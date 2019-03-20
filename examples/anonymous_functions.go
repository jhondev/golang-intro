package main

import "fmt"

var f = func() {
	fmt.Println("test anonymous")
}

func main() {
	f()

	f2 := func(textparam string) {
		println(textparam)
	}

	f2("internal scope")

	func() {
        fmt.Println("Functions anonymous")
    }()
}
