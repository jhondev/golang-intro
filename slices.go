package main

func main() {
	planets := []string{"Earth", "Mars", "Jupiter"}
	planetsArray := [3]string{"Earth", "Mars", "Jupiter"}
	testSlices(planets)
	testArray(planetsArray)
	println(len(planets))
	println(len(planetsArray))
}

func testSlices(slice []string) {
	// slice = append(slice, "next planet", "n2", "n3", "n4", "n5")
	for _, item := range slice {
		println(item)
	}
	println("slice capacity")
	println(cap(slice))
}

func testArray(array [3]string) {
	// error: array = append(array, "next planet in array")
	for _, item := range array {
		println(item)
	}
}
