package insertinsort

import (
	"fmt"
	insertinsort "insertionsort"
	"testing"
)

func TestInsertionsort(t *testing.T) {
	input := []int{5, 4, 3, 2, 1}
	fmt.Println(input)
	testarray := make([]int, 5)
	copy(testarray, insertinsort.Insertinsort(input))
	want := [5]int{1, 2, 3, 4, 5}
	for i := 0; i <= 4; i++ {
		if testarray[i] != want[i] {
			t.Errorf("expected:%v,got:%v", want[i], testarray[i])
		}
	}
}

func BenchmarkInsertion(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	x := make([]int, 5)
	for i := 0; i < b.N; i++ {
		copy(input, x)
		insertionsort.Insertinsort(x)
	}
}
