package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	for index := 0; index < 5; index++ {
		go sleepyGopher(index, c)		
	}
	for index := 0; index < 5; index++ {
		gopherID := <- c

		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore...")
	c <- id
}