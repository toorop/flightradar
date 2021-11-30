package flightradar

type Aircraft struct {
	Fr24Id           string  `json:"fr_24_id"`
	IcaoRegistration string  `json:"icao_registration"` // ICAO 24-bit address The ICAO24 code (sometimes called the Mode S code) is a 24-bit unique number that is assigned to each vehicle or object that can transmit ADS-B messages. It is usually transmitted by aircraft but some airport ground vehicles and multilateration towers also have ICAO24 codes assigned to them.
	Latitude         float64 `json:"lat"`               // latitude
	Longitude        float64 `json:"lon"`               // longitude
	Bearing          uint8   `json:"heading"`           // heading in degrees
	Altitude         uint    `json:"altitude"`          // altitude in feet
	Speed            uint    `json:"speed"`             // speed in knots
	SquawkCode       string  `json:"squawk_code"`       // squawk code
	RadarID          string  `json:"radar_id"`          // radar ID
	IcaoModel        string  `json:"icao_model"`        // ICAO model type
	Registration     string  `json:"registration"`      // registration
	Timestamp        int64   `json:"timestamp"`         // timestamp
	Origin           string  `json:"origin"`            // origin airport IATA code
	Destination      string  `json:"destination"`       // destination airport IATA code
	FlightNumber     string  `json:"flight_number"`     // flight number
	IsOnGround       bool    `json:"on_ground"`         // is on ground
	RateOfClimb      uint    `json:"rate_of_climb"`     // rate of climb in feet per minute
	CallSign         string  `json:"callsign"`          // callsign
	IsGlider         bool    `json:"is_glider"`         // is glider
	Company          string  `json:"company"`           // company
}
