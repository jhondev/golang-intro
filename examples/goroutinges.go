package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepyGopher()
	for index := 0; index < 4; index++ {
		time.Sleep(1 * time.Second)
		println("waiting")	
	}
	
}

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...")
}
