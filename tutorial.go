package main

import "fmt"

func main() {
	name := "Вася"
	fmt.Println(name)

	setName(name)
	fmt.Println(name)
}

func setName(name string) {
	name = "Петя"
}
