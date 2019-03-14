package main

func main() {
	funcText := myText("testing funcs")
	println(funcText)
}

func myText(text string) string {
	return "hey " + text
}
