package main

import "log"

func main() {
	Foot := 0.3048
	Meter := 1.000
	footPerMeter := Meter / Foot
	log.Println(footPerMeter)
}
