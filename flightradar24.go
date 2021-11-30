package flightradar

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
)

type RadarOptions struct {
	Faa               bool // use US/canada data source
	Flarm             bool // use Flarm data source
	Mlat              bool // use MLAT data source
	Adsb              bool // use ADS-B data source
	InAir             bool // get in-air aircraft
	OnGround          bool // get on-ground aircraft
	Inactive          bool // get inactive aircraft (on ground ;) )
	Gliders           bool // get gliders
	EstimatedPosition bool // get estimated position
}

type RadarResponse struct {
	FullCount int `json:"full_count"` // total number of aircraft
	Version   int `json:"version"`    // version of the data
	Aircrafts []Aircraft
}

type FRadar24 struct{}

func (r FRadar24) Scan() ([]Aircraft, error) {

	resp, err := http.Get("https://data-live.flightradar24.com/zones/fcgi/feed.js?faa=1&bounds=49.072%2C48.918%2C2.431%2C2.695&satellite=1&mlat=1&flarm=1&adsb=1&gnd=1&air=1&vehicles=1&estimated=1&maxage=14400&gliders=1&stats=1")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//log.Println(string(body))

	var response map[string]interface{}

	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	rr := RadarResponse{}
	//rr.Aircrafts = make(map[string]interface{})

	for k, v := range response {
		switch k {
		case "full_count":
			rr.FullCount = int(v.(float64))
		case "version":
			rr.Version = int(v.(float64))
		case "stats":
			//
		default:
			vMap := v.([]interface{})
			aircraft := Aircraft{}
			aircraft.Fr24Id = k
			aircraft.IcaoRegistration = vMap[0].(string)
			aircraft.Latitude = vMap[1].(float64)
			aircraft.Longitude = vMap[2].(float64)
			aircraft.Bearing = uint8(vMap[3].(float64))
			aircraft.Altitude = getUint(vMap[4])
			aircraft.Speed = uint(vMap[5].(float64))
			aircraft.SquawkCode = getString(vMap[6])
			aircraft.RadarID = getString(vMap[7])
			aircraft.Registration = getString(vMap[8])
			aircraft.IcaoModel = getString(vMap[9])
			aircraft.Timestamp = getInt64(vMap[10])
			aircraft.Origin = getString(vMap[11])
			aircraft.Destination = getString(vMap[12])
			aircraft.FlightNumber = getString(vMap[13])
			aircraft.IsOnGround = getBoolean(vMap[14])
			aircraft.RateOfClimb = getUint(vMap[15])
			aircraft.CallSign = getString(vMap[16])
			aircraft.IsGlider = getBoolean(vMap[17])
			aircraft.Company = getString(vMap[18])

			// append
			rr.Aircrafts = append(rr.Aircrafts, aircraft)
		}
	}

	return rr.Aircrafts, err
}

func getString(v interface{}) string {
	if reflect.TypeOf(v).String() == "string" {
		return v.(string)
	}
	return ""
}

func getInt64(v interface{}) int64 {
	if reflect.TypeOf(v).String() == "int64" {
		return v.(int64)
	}
	return 0
}

func getBoolean(v interface{}) bool {
	return v.(float64) == 1
}

func getUint(v interface{}) uint {
	if reflect.TypeOf(v).String() == "uint" {
		return v.(uint)
	}
	return 0
}
