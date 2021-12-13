package main

import "fmt"

var a, b, c int //Глобальные переменные

/*func main() {
	//message := "Я скоро стану GOLANG NINJA"
	//var message string = "Я скоро стану GOLANG NINJA"
	//message := "Я скоро стану GOLANG NINJA" // Обявления и инецилизация переменной
	//var message int // Обявления переменной
	//message = 12 // Инецилизация переменной
	//fmt.Println(reflect.TypeOf(message))
	//var message string
	var number float64
	var b bool

	b = true
	number = 12.4 // Локальные переменные

	message := []byte("asd")

	fmt.Println(message)
	fmt.Println(number)
	fmt.Println(b)
}*/

func main() {
	//var message string
	//message = sayHello("Максим", 21)
	//printMessage(message)
	/*message, err := enterTheClub(55)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(message)

		fmt.Println(prediction("чт"))
	}*/

	/*func printMessage(message string) {
		fmt.Println(message)
	}*/

	/*func print() {
	  	//a, b, c := 4, 5, 6
	  	fmt.Println(a, b, c)
	  }

	  func sayHello(name string, age int) string {
	  	return fmt.Sprintf("Привет, %s! Тебе %d лет", name, age)
	  }

	  func enterTheClub(age int) (string, error) {
	  	if age >= 18 && age < 45 {
	  		return "Входи, только акуратно", nil
	  	} else if age >= 45 && age < 65 {
	  		return "Вам точно понравиться это музыка?", nil
	  	} else if age >= 65 {
	  		return "Это уже слишком для вас", errors.New("you are too old")
	  	} else {
	  		return "Тебе не 18-ти", errors.New("you are too young")
	  	}
	  }

	  func prediction(dayOfWeek string) (string, error) {
	  	switch dayOfWeek {
	  	case "пн":
	  		return "Хорошего тебе недели !", nil
	  	case "вт":
	  		return "Прекрасного тебе вторника!", nil
	  	case "ср":
	  		return "Замечательной тебе среды!", nil
	  	case "чт":
	  		return "Четверг - это маленькая пятница! Только не переборщи!!!", nil
	  	default:
	  		return "Невалидный день недели!!", errors.New("invalid day of the week")
	  	}*/
	fmt.Println(findMin(1, 2, 568, 234864, -56464, -5648, 586))
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}
