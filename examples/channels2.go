package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)

	for index := 0; index < 5; index++ {
		go sleepyGopher(index, c)
	}
	timeout := time.After(2 * time.Second)
	for index := 0; index < 5; index++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout:
			fmt.Println("My patience ran out")
			return
		}
	}
}

func sleepyGopher(id int, c chan int) {
	// time.Sleep(3 * time.Second)
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("... ", id, " snore...")
	c <- id
}
