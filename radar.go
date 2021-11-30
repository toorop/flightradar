package flightradar

type Radar interface {
	Scan() ([]Aircraft, error)
}
