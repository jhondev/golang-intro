package main

func main() {
	var k kelvin = 294
	c := k.celsius()

	println(c)
}

// custom type
type kelvin float64
type celsius float64

// method
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
