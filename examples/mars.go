// My weight loss program.
package main

import "fmt"
import "math/rand"
import "strings"
import "strconv"

func main() {
	fmt.Printf("My weight on the surface of Mars is %v lbs, %v", 149.0*0.3783, 59)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

	var onev = 1
	onev = 3
	println(onev)
	// var onev1, onev2, one3 = 1, 2, "test"
	println("**********shorcuts")
	var weight = 1
	println(weight)
	weight = weight * 2
	println(weight)
	weight *= 6
	println(weight)
	weight++
	println(weight)

	var num = rand.Intn(10) + 1
	fmt.Println(num)

	var contain = strings.Contains("tendador", "dado")
	fmt.Printf("value: %v \n", contain)

	if contain {
		println("it contains")
	} else {
		println("no contains")
	}

	switch "arriba" {
	case "arriba":
		println("arriba1")
	case "dddd":
	case "no arriba":
		println("no arriba 1")
	default:
		println("default")
	}

	var testScope = "scope1"
	var count = 3
	for count > 0 {
		println(testScope)
		var testScope = "scope2"
		println(testScope)
		count--
		fmt.Printf("repeat: %v \n", count)
	}

	// short declaration
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}

	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}

	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)

	grade := 'A'
	fmt.Printf("%c = %[1]v \n", grade)

	peace := "shalom"
	peace = "salām"
	println(peace)

	println("type conversion " + strconv.Itoa(25))
}
