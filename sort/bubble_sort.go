package sort

import "fmt"

func bubble_sort(list []int) []int {
	n := len(list)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

func BubbleSort() {
	unsorted_list := []int{40, 28, 3, 12, 29, 47, 41, 11, 32, 36, 22, 42, 25, 7, 10, 34, 15, 2, 39, 24, 5, 14, 13, 30, 31, 27, 46, 8, 9, 4, 19, 50, 16, 18, 6, 38, 35, 23, 33, 45, 20, 17, 43, 37, 48, 1, 26, 21, 49, 44}
	
	sorted := bubble_sort(unsorted_list)
	fmt.Println(sorted)
}
