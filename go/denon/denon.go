package denon

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/huin/goupnp/dcps/av1"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger
var logger *zap.SugaredLogger

func init() {
	zapLogger = zap.NewExample()
	logger = zapLogger.Sugar()
}

// Denon is an interface to a networked Denon AVR
type Denon struct {
	hostname string

	queryTicker *time.Ticker

	status     Status
	statusTime time.Time
}

// Input describes an A/V input on the Denon AVR such as GAME, SAT/CBL, DVD, etc.
type Input struct {
	ID   string `xml:"index,attr"`
	Name string `xml:",chardata"`
}

// DenonZone blah
type zoneResponse struct {
	Power     string  `xml:"Power>value"`
	InputName string  `xml:"InputFuncSelect>value"`
	Volume    string  `xml:"MasterVolume>value"`
	Inputs    []Input `xml:"VideoSelectLists>value"`
}

// Status contains the status read from a Denon AVR.
type Status struct {
	// Power is true if the device is powered on (not on standby)
	Power bool

	// Input is the name of the current A/V input
	Input Input

	// Volume is the sound volume. Zero is the lowest value. 80 is max volume.
	Volume float32

	// Inputs is a map of Display -> ID for A/V inputs.
	// The named inputs like "GAME" or "DVD" never change internally even
	// if you change the display name.
	Inputs []Input `xml:"VideoSelectLists>value"`
}

// URL reports the url for this Denon device.
func (denon *Denon) URL() *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   denon.hostname,
		Path:   "/goform/formMainZone_MainZoneXml.xml",
	}
}

// Query asks the device for the current status (input, volume, power status)
func (denon *Denon) Query() (status Status, err error) {
	endpoint := denon.URL()
	response, err := http.Get(endpoint.String())

	if err != nil {
		logger.Errorw("Failed sending command to Denon AVR", "endpoint", endpoint, "error", err)
		return
	}

	defer response.Body.Close()

	zr := zoneResponse{}
	decoder := xml.NewDecoder(response.Body)
	err = decoder.Decode(&zr)
	if err != nil {
		logger.Errorw("Failure reading response from Denon AVR", "endpoint", endpoint, "error", err)
		return
	}

	log.Printf("Zone: %#v", zr)
	err = processQuery(zr, &status)

	denon.status = status
	denon.statusTime = time.Now()
	return
}

func processQuery(zr zoneResponse, status *Status) (err error) {
	if zr.Volume == "--" {
		status.Volume = 0
	} else {
		var i float64
		i, err = strconv.ParseFloat(zr.Volume, 32)
		if err != nil {
			logger.Warnw("Unexpected master volume value?", "value", zr.Volume, "error", err)
			return
		}
		status.Volume = convertVolume(float32(i))
	}

	status.Power = zr.Power == "ON"

	status.Inputs = make([]Input, 0)
	for _, input := range zr.Inputs {
		input.ID = strings.TrimSpace(input.ID)
		input.Name = strings.TrimSpace(input.Name)

		if zr.InputName == input.Name {
			status.Input = input
		}
		status.Inputs = append(status.Inputs, input)
	}

	return
}

func convertVolume(value float32) float32 {
	// On my Denon AVR, when the volume is set minimum (1) the value in the API appears as -79
	// zero I assume is -80. When I call MV45 (set master volume to 45) the volume in the api
	// shows up as -35.
	// Therefore, the "zero" value for volume is -80. Offset all volumes by 80 to scale them
	// based on zero.
	return value + 80
}

// Command sends a Denon AVR command. These commands are documented in the Denon AVR control protocol.
// You can usually find the PDF for the control protocol online.
func (denon *Denon) Command(command string) {
	endpoint := url.URL{
		Scheme:   "http",
		Host:     denon.hostname,
		Path:     "/goform/formiPhoneAppDirect.xml",
		RawQuery: command,
	}

	_, err := http.Get(endpoint.String())

	if err != nil {
		logger.Errorw("Failed sending command to Denon AVR", "command", command, "endpoint", endpoint)
		return
	}
}

func New(hostname string) *Denon {
	return &Denon{hostname: hostname}
}

// Discover uses network discovery find a neighboring Denon AVR devices.
func Discover() (denon *Denon) {
	clients, errs, err := av1.NewAVTransport1Clients()
	if len(errs) > 0 {
		for _, e := range errs {
			logger.Infow("Error", "error", e)
		}
	}
	if err != nil {
		logger.Infow("Error", "error", err)
	}

	if len(clients) == 0 {
		logger.Infow("no Av1 clients found on network?")
		return
	}

	for _, c := range clients {
		logger.Debugw("Found Av1 Client", "device", c.RootDevice.Device.Manufacturer)
		if c.RootDevice.Device.Manufacturer == "Denon" {
			denon = &Denon{hostname: c.Location.Hostname()}
			logger.Infow("Discovered Denon AVR", "address", denon.hostname, "device", c.RootDevice.Device.FriendlyName)
			return
		}
	}

	logger.Warn("No denon found")
	return
}
