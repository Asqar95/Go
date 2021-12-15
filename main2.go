package main

import (
	"errors"
	"fmt"
)

func main() {
	/*messages := make([]string, 5)
	messages = append(messages, "6")
	messages = append(messages, "7")
	messages = append(messages, "8")
	messages = append(messages, "9")
	messages = append(messages, "10")
	fmt.Println(messages)
	fmt.Println(len(messages))
	fmt.Println(cap(messages))*/
	//messages := []string{"1", "2", "3"}

	/*matrix := make([][]int, 10)

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			matrix[y] = make([]int, 10)
			matrix[y][x] = x
		}
		fmt.Println(matrix[x])
	}*/

	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
}

func printMessage(messages []string) error {
	if len(messages) == 0 {
		return errors.New("employ array")
	}

	messages[1] = "5"

	fmt.Println(messages)

	return nil
}
