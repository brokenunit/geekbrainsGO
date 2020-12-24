package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("введите число разрядов")
	var m string // в коментарии к ПР, вы предложили использовать вместо фиксированного числа разрядов переменную
	var i int
	i, err := strconv.Atoi(m)
	if err != nil {
		fmt.Println("101")
	}
	arr := []int{}
	//цикл добавляет в "массив" элементы из os.Stdin по одному, с bufio разобраться не успел
	for n := 0; n < i; n++ {
		var inputNew string
		var numberNew int
		fmt.Println("введите элемент массива")
		fmt.Fscan(os.Stdin, &inputNew)
		numberNew, err := strconv.Atoi(inputNew)
		if err != nil {
			fmt.Println("110")
			os.Exit(1)
		}
		arr = append(arr, numberNew)
	}
	fmt.Println(arr)
	ar := make([]int, i)
	copy(ar, arr)
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ { // тут было замечание по поводу лишних проходов в цикле, возможно теперь стало лучше
			if ar[j] > ar[i] {
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}
	fmt.Println(ar)
}
