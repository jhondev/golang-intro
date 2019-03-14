package main

import "fmt"

func main() {
	type location struct {
		lat  float64
		long float64
	}
	curiosity := location{long: 45, lat: 1}
	curiosity2 := curiosity
	curiosity2.lat = 3
	var curiosity3 location
	fmt.Println(curiosity)
	fmt.Println(curiosity2)
	fmt.Println(curiosity3)
}
