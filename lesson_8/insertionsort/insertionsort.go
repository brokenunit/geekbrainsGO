package insertinsort

import (
	"fmt"
)

func Insertinsort(arr []int) []int {
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
	return ar
}
