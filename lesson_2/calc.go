package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var opOne, opTwo, res float64
	// ввод операда 1
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&opOne)
	//fmt.Println(opOne)
	// ввод операнда 2
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&opTwo)
	//fmt.Println(opTwo)
	// действие над числами
	var op string
	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = opOne + opTwo
	case "-":
		res = opOne - opTwo
	case "*":
		res = opOne * opTwo
	case "/":
		if opTwo == 0 {
			os.Exit(100)
		}
		res = opOne / opTwo
	case "^":
		res = math.Pow(opOne, opTwo)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Println(opOne, "+", opTwo, "=", res)

}
