package main

import "fmt"

func main() {
	var num int = 250
	var word string = "Apple"
	var minimunWage int = 10
	var workingHour int = 20

	var income int = minimunWage * workingHour

	num = 10

	fmt.Println(minimunWage, workingHour, income)
	fmt.Println(word)
	fmt.Println(num)

}
