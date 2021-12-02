package main

import (
	"fmt"

	"github.com/toorop/flightradar"
)

func main() {

	options := flightradar.RadarOptions{
		Bounds:            [4]float64{43.65813137907139, 1.3274154066556807, 43.60904450519146, 1.4013737722054362},
		Faa:               true,
		Flarm:             true,
		Mlat:              true,
		Adsb:              true,
		InAir:             true,
		OnGround:          true,
		Inactive:          true,
		Gliders:           true,
		EstimatedPosition: true,
	}

	radar, err := flightradar.NewFRadar24(options)
	if err != nil {
		panic(err)
	}
	aircrafts, err := radar.Scan()
	if err != nil {
		panic(err)
	}
	for _, aircraft := range aircrafts {

		fmt.Printf("%v\n", aircraft)
	}

}
