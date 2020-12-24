package main

import (
	"fmt"
	"os"
	"strconv"
)

//функция для вичисления числа фиббоначи
func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n < 0 {
		fmt.Println("неверный индекс")
		os.Exit(1)
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {
	fmt.Println("введите индекс числа")
	var inputNum string
	var n int
	fmt.Scanln(&inputNum)
	n, err := strconv.Atoi(inputNum)
	if err != nil {
		fmt.Println("110")
		os.Exit(1)
	}
	fmt.Println("Число Фиббоначи с индексом", n, "равно", fibo(n))
}
