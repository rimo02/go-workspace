package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// every package can have only 1 main func
func main() {
	var username string = "Rimo Ghosh"
	fmt.Println("Hello " + username)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name of student:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Your name is ", input)

	fmt.Println((time.Now()))

	num := 25
	var ptr = &num
	fmt.Println("Address= ", ptr)

}
