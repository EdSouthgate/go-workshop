package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("Calculator Go 1.0")
	fmt.Println("=================")
	fmt.Println("Which calculation do you want to perform? (add, subtract, multiply, divide)")
	fmt.Scanf("%s\n", &operation) // & required as we wish to modify the value by default a copy would be passed so we pass a pointer to the position in memory. 
	fmt.Println("Enter first number")
	fmt.Scanf("%d\n", &number1) // \n required on windows for reasons I don't exactly understand yet but it was the fix to it not working as per the demo
	fmt.Println("Enter second number")
	fmt.Scanf("%d\n", &number2)
	switch operation {
		case "add": 
			fmt.Println(number1 + number2)
			case "subtract": 
			fmt.Println(number1 - number2)
		case "multiply":
			fmt.Println(number1 * number2)
		case "divide":
			fmt.Println(number1 / number2)

	}

	
}
