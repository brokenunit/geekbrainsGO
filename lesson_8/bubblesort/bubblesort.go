package bubblesort

//Функция сортировки "пузырьком"
func Bubblesort(arr []int) []int {
	ar := make([]int, 5)
	copy(ar, arr)
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			if ar[j] > ar[i] {
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}
	// ar[4] = 10 //не придуамл как еще вызвать ошибку
	return ar
}
