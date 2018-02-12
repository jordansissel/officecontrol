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
	hostname string

	queryTicker *time.Ticker
}

type DenonStatus struct {
	MasterVolume float32
}

func (denon *Denon) Query() (DenonStatus, error) {
	endpoint := url.URL{
		Scheme: "http",
		Host:   denon.hostname,
		Path:   "/goform/formMainZone_MainZoneXml.xml",
	}

	response, err := http.Get(endpoint.String())

	if err != nil {
		logger.Errorw("Failed sending command to Denon AVR", "endpoint", endpoint)
		return DenonStatus{}, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		logger.Errorw("Failure reading response from Denon AVR", "endpoint", endpoint, "error", err)
		return DenonStatus{}, err
	}

	item := DenonItem{}
	xml.Unmarshal(body, &item)
	return processQuery(&item), nil
}

func processQuery(item *DenonItem) (result DenonStatus) {
	text := item.DenonMasterVolume.DenonValue[0].Text
	if text == "--" {
		result.MasterVolume = 0
	} else {
		i, err := strconv.ParseFloat(text, 32)
		if err != nil {
			logger.Warnw("Unexpected master volume value?", "value", text, "error", err)
		} else {
			result.MasterVolume = convertVolume(float32(i))
		}
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

func Discover() (denon *Denon) {
	clients, _, _ := av1.NewAVTransport1Clients()

	// XXX: Verify if the AV1 location is a valid Denon AVR.
	// XXX: This can be done by fetching the Location (description.xml) and looking for AVR in the modelName.
	if len(clients) == 0 {
		fmt.Println("no Av1 clients found on network?")
		return
	}

	avr := clients[0]

	denon = &Denon{hostname: avr.Location.Hostname()}
	logger.Infow("Discovered Denon AVR", "address", denon.hostname)
	return
}
