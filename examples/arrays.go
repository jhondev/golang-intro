package main

func main() {
	var planets [8]string

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]
	println(earth)
	println(len(planets))

	// out of range err
	// i := 11
	// println(planets[i])

	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	for index, dwarf := range dwarfs {
		println(index, dwarf)
	}
}
