package main

func main() {
	 response := defer test()
	println("hey don't panic")
	println(response)
}

func test() int {
	zz := 0
	re := 1 / zz
	return re
}
