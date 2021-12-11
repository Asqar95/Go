package main

import ("fmt")

func main() {
	var name string
	var age int
	var weight float32

	name = "Максим"
	weight = 69.3
	age = 28
	printPersonInfo(name, age, weight)
	printPersonInfo(name: "Вася" , age "33", weight"78.3")
}
func printPersonInfo(name string, age int, age int, weight float32) {
	fmt.Printf(format:"Имя: %s\nВозрст: %d/nВес: %f/n", name, age, weight)
}
