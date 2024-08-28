package algods

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var mid int
	var left, right, divleft, divright []int
	mid = len(arr) / 2
	left = arr[:mid]
	right = arr[mid:]
	divleft = MergeSort(left)
	divright = MergeSort(right)
	output := merge(divleft, divright)

	return output
}

func merge(left []int, right []int) []int {
	var output []int
	var l, r int
	l = 0
	r = 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			output = append(output, left[l])
			l++
		} else {
			output = append(output, right[r])
			r++
		}
	}
	fmt.Println(output)
	return output
}
