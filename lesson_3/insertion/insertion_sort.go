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
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
	fmt.Println(ar)
}
