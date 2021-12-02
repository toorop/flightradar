package flightradar

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const FRadar24BaseURL = "https://data-live.flightradar24.com/zones/fcgi/feed.js?"

type FRadar24 struct {
	options RadarOptions
}

func NewFRadar24(options RadarOptions) (*FRadar24, error) {
	// todo: validate options
	return &FRadar24{
		options: options,
	}, nil
}

// Scan scans the radar for flights
// https://data-live.flightradar24.com/zones/fcgi/feed.js?
//faa=1&
//bounds=49.072%2C48.918%2C2.431%2C2.695&  49.072,48.918,2.431,2.695
//satellite=1&
//mlat=1&
//flarm=1&
//adsb=1&
//gnd=1&
//air=1&
//vehicles=1&
//estimated=1&
//maxage=14400&
//gliders=1&stats=1

func (r FRadar24) Scan() ([]Aircraft, error) {

	//var fullCount, version int

	// URL
	url := FRadar24BaseURL

	// faa
	if r.options.Faa {
		url += "faa=1&"
	}
	// bounds
	url += fmt.Sprintf("bounds=%f,%f,%f,%f&", r.options.Bounds[0], r.options.Bounds[2], r.options.Bounds[1], r.options.Bounds[3])

	//satellite
	/*if r.options.satellite {
		url += "satellite=1&"
	}

	*/

	//mlat
	if r.options.Mlat {
		url += "mlat=1&"
	}
	// flarm
	if r.options.Flarm {
		url += "flarm=1&"
	}
	// adsb
	if r.options.Adsb {
		url += "adsb=1&"
	}
	// gnd
	if r.options.OnGround {
		url += "gnd=1&"
	}
	// air
	if r.options.InAir {
		url += "air=1&"
	}
	// vehicles
	if r.options.Inactive {
		url += "vehicles=1&"
	}
	// estimated
	url += "estimated=1&"

	// max age
	//url += fmt.Sprintf("maxage=%d&", r.options.maxage)
	// gliders
	if r.options.Gliders {
		url += "gliders=1"
	}

	fmt.Printf("URL: %s\n", url)

	resp, err := http.Get(url)
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

	//rr := RadarResponse{}
	aircrafts := make([]Aircraft, 0)
	//rr.Aircrafts = make(map[string]interface{})

	for k, v := range response {
		switch k {
		case "full_count":
			//fullCount = int(v.(float64))
		case "version":
			//version = int(v.(float64))
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
			aircrafts = append(aircrafts, aircraft)
		}
	}

	return aircrafts, err
}
