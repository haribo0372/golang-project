package main

import (
	"errors"
	"fmt"
)

func main() {
	var name = "Johan"
	fmt.Println("Hello, World!")
	hello(name)
	printEven(0, 8)
	apply(9, 8, "+")
	apply(9, 8, "-")
	apply(9, 8, "*")
	apply(9, 8, "/")
	var _, error1 = apply(9, 0, "/")
	var _, error2 = apply(9, 8, "&")

	if error1 != nil {
		fmt.Println(error1.Error())
	}

	if error2 != nil {
		fmt.Println(error2.Error())
	}
}

func printEven(a, b int64) {
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println("Четное число ", i)
		}
	}
}

func hello(name string) {
	fmt.Println("Hello, " + name)
}

func apply(a float64, b float64, operation string) (float64, error) {
	var template = "a : %.2f, b: %.2f operation: %s, result: %.2f\n"
	var result = 0.0
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return 0.0, errors.New("CE: Деление на ноль")
		}
		result = a / b
	default:
		return 0.0, errors.New("CE: Операция " + operation + " не поддерживается")
	}

	fmt.Printf(template, a, b, operation, result)
	return result, nil
}
