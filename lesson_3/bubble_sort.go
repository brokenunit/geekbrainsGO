package main

import (
	"fmt"
	"os"
)

func main() {
	arr := []int{}
	//цикл добавляет в "массив" элементы из os.Stdin по одному, с bufio разобраться не успел
	for n := 0; n < 5; n++ {
		var numberNew int
		fmt.Println("введите элемент массива")
		fmt.Fscan(os.Stdin, &numberNew) // здесь должна была быть проверка типа введенных данных, эхх
		arr = append(arr, numberNew)
	}
	fmt.Println(arr)
	ar := make([]int, 5)
	copy(ar, arr)
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			if ar[j] > ar[i] {
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}
	fmt.Println(ar)
}
