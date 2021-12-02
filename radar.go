package flightradar

type RadarOptions struct {
	Bounds            [4]float64 // [upperLeftLat, upperLefLng, lowerRightLat, lowerRightLng]
	Faa               bool       // use US/canada data source
	Flarm             bool       // use Flarm data source
	Mlat              bool       // use MLAT data source
	Adsb              bool       // use ADS-B data source
	InAir             bool       // get in-air aircraft
	OnGround          bool       // get on-ground aircraft
	Inactive          bool       // get inactive aircraft (on ground ;) )
	Gliders           bool       // get gliders
	EstimatedPosition bool       // get estimated position
}

type Radar interface {
	Scan() ([]Aircraft, error)
}
