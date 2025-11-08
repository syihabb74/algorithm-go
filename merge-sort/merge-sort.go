package main

import (
	"fmt"
)

func merge(unsorted_arr []int) []int {
	if len(unsorted_arr) < 2 {
		return unsorted_arr
	}
	var mid int = len(unsorted_arr) / 2
	return mergeSort(merge(unsorted_arr[:mid]), merge(unsorted_arr[mid:]))
}
// return new array but not shorting in place
func mergeSort(left []int, right []int) []int {
	var result []int = make([]int, 0, len(left) + len(right))
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {

			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r]);
			r++
		}
	}

	if l < len(left) {
			
			result = append(result, left[l:]...)
	}
	if r < len(right) {
			result = append(result, right[r:]...)
	}
	return result

}

func main() {

	var unsorted_arr []int = []int{
		1, 3, 2, 4, 7, 9,
	}

	fmt.Println(merge(unsorted_arr)) // return new array and shorting from existing slice 
	fmt.Println(unsorted_arr) // still unsorted


}
