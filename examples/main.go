package main

import (
	"fmt"

	"github.com/toorop/flightradar"
)

func main() {
	radar, err := flightradar.NewFlightRadar24()
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
