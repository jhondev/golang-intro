package main

import "fmt"

func main() {
	bestFinish := bestLeagueFinishes(12, 13, 8, 9, 10)
	fmt.Println(bestFinish)
}

func bestLeagueFinishes(finishes ...int) int {
	return 2
}
