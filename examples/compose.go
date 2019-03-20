package main

import "fmt"

type celsius float64

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type report struct {
	sol         int
	temperature temperature
	location    location
}

func (t temperature) average() celsius {
    return (t.high + t.low) / 2
}

func (r report) average() celsius {
    return r.temperature.average()
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report)

	fmt.Printf("a balmy %vº C\n", report.temperature.high)

	fmt.Printf("average %v° C\n", t.average())

	fmt.Printf("average from report %v° C\n", report.average())
}
