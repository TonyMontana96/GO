package main

import "log"

func main() {
	var x [6]float64
	x[0] = 111
	x[1] = 90
	x[2] = 80
	x[3] = 44
	x[4] = 22
	x[5] = 11

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	log.Println(total / float64(len(x)))
}
