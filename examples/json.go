package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Lat float64 `json:"lat"`
		Long float64 `json:"long"`
	}
	curiosity := location{Lat: 25.3, Long: 36.7}
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
