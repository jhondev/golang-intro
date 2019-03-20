package main

func main() {
	answer := 42
	println(&answer)
	addresspointer := &answer
	println(*addresspointer)
	
	println("*********struct")
	person1 := person{age: 10}
	birthdayStruct(person1)
	println(person1.age)
	
	println("********struct with pointer")
	person2 := person{age:15}
	birthdayPointer(&person2)
	println(person2.age)

	println("***method receiver struct")
	person1.birthdayStruct()
	println(person1.age)

	println("***method receiver pointer")
	person2.birthdayPointer()
	println(person2.age)

	value := 4
	test(&value)
	println(value)

	value2 := 4
	test2(value2)
	println(value2)
}

type person struct{ age int }

func birthdayStruct(person person) {
	person.age++
	println(person.age)
}

func birthdayPointer(person *person) {
	person.age++
	println(person.age)
}

func (p person) birthdayStruct() {
	p.age++	
}

func (p *person) birthdayPointer() {
	p.age++
}

func test(value *int) {
	(*value)++;
}

func test2(value int) {
	value++;
}

