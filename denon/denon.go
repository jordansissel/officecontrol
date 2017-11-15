package denon

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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
	hostname     string
	masterVolume float32

	queryTicker *time.Ticker
}

func (denon *Denon) Query() (err error) {
	endpoint := url.URL{
		Scheme: "http",
		Host:   denon.hostname,
		Path:   "/goform/formMainZone_MainZoneXml.xml",
	}

	response, err := http.Get(endpoint.String())

	if err != nil {
		logger.Errorw("Failed sending command to Denon AVR", "endpoint", endpoint)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		logger.Errorw("Failure reading response from Denon AVR", "endpoint", endpoint, "error", err)
		return
	}

	item := DenonItem{}
	xml.Unmarshal(body, &item)
	denon.processQuery(&item)

	return
}

func (denon *Denon) processQuery(item *DenonItem) {
	text := item.DenonMasterVolume.DenonValue[0].Text
	if text == "--" {
		denon.masterVolume = 0
	} else {
		i, err := strconv.ParseFloat(text, 32)
		if err != nil {
			logger.Warnw("Unexpected master volume value?", "value", text, "error", err)
		} else {
			denon.masterVolume = convertVolume(float32(i))
		}
	}
}

func convertVolume(value float32) float32 {
	// On my Denon AVR, when the volume is set minimum (1) the value in the API appears as -79
	// zero I assume is -80. When I call MV45 (set master volume to 45) the volume in the api
	// shows up as -35.
	// Therefore, the "zero" value for volume is -80. Offset all volumes by 80 to scale them
	// based on zero.
	return value + 80
}

// Send a raw Denon AVR command. These commands are documented in the Denon AVR control protocol.
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

func (denon *Denon) watch() {
	if denon.queryTicker != nil {
		denon.queryTicker.Stop()
		denon.queryTicker = nil
	}

	denon.queryTicker = time.NewTicker(time.Second * 2)
	go denon.observe()
}

func (denon *Denon) observe() {
	for range denon.queryTicker.C {
		err := denon.Query()
		if err != nil {
			logger.Error("Failed querying Denon AVR")
		}
	}
}

func Discover() (denon *Denon) {
	clients, _, _ := av1.NewAVTransport1Clients()

	// XXX: Verify if the AV1 location is a valid Denon AVR.
	// XXX: This can be done by fetching the Location (description.xml) and looking for AVR in the modelName.
	if len(clients) == 0 {
		fmt.Println("no Av1 clients found on network?")
		return
	}

	avr := clients[0]
	fmt.Printf("r: %#v\n", avr.Location.Hostname())

	denon = &Denon{hostname: avr.Location.Hostname()}

	denon.watch()
	return
}
