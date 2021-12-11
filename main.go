package main

import (
	"fmt"
	"reflect"
)

func main() {
	message := "Я скоро стану Golang Ninja"

	fmt.Println(reflect.TypeOf(message))
}
