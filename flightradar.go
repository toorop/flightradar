package flightradar

func NewFlightRadar24() (Radar, error) {
	radar := new(FRadar24)
	return radar, nil
}
