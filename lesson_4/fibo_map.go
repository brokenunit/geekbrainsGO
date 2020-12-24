package main

import (
	"fmt"
	"os"
	"strconv"
)

func fiboMap(n int) int {
	// m := make(map[int]int)
	m := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}
	var j int
	if n == 2 {
		return 1
	}
	if n > 93 {
		fmt.Println("неверный индекс")
		os.Exit(1)
	}
	if n < 0 {
		fmt.Println("неверный индекс")
		os.Exit(1)
	}
	for i := 2; i < n; i++ {
		m[i] = m[i-1] + m[i-2]
		j = i
	}
	return m[j]
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
	fmt.Println("Число Фиббоначи с индексом", n, "равно", fiboMap(n))
}
