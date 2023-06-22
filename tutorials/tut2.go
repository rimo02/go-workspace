package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	var arr [4]string
	arr[0] = "Apple"
	arr[1] = "Mango"
	// arr[2]="Banana"
	arr[3] = "Dustbin"
	fmt.Println(arr)

	scores := make([]int, 3)
	scores[0] = 10
	scores[1] = 20
	scores[2] = 30

	scores = append(scores, 30, 40, 50)
	fmt.Println(scores)

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	// delete(languages,"RB")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("Key: %v , Value:%v \n", key, value)
	}

	rimo := User{"Rimo", 21}
	fmt.Println(rimo)

	if 8%4 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	for num := 0; num < len(scores); num++ {
		fmt.Println(scores[num])
	}

	sum := greeter(10, 20)
	fmt.Println("Sum =", sum)
	fmt.Println("Sum =", Adder(10, 20, 30, 40, 50))

}
func greeter(num1 int, num2 int) int {
	return num1 + num2

}

func Adder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total

}
