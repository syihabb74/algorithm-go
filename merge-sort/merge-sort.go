package main

import (
	"fmt"
)

// func merge(unsorted_arr []int) []int {
// 	if len(unsorted_arr) < 2 {
// 		return unsorted_arr
// 	}
// 	var mid int = len(unsorted_arr) / 2
// 	return mergeSort(merge(unsorted_arr[:mid]), merge(unsorted_arr[mid:]))
// }
// // return new array but not shorting in place
// func mergeSort(left []int, right []int) []int {
// 	var result []int = make([]int, 0, len(left) + len(right))
// 	l := 0
// 	r := 0
// 	for l < len(left) && r < len(right) {
// 		if left[l] < right[r] {

// 			result = append(result, left[l])
// 			l++
// 		} else {
// 			result = append(result, right[r]);
// 			r++
// 		}
// 	}

// 	if l < len(left) {
			
// 			result = append(result, left[l:]...)
// 	}
// 	if r < len(right) {
// 			result = append(result, right[r:]...)
// 	}
// 	return result

// }


func merge_in_place (unsorted []int) []int {

	// fmt.Println(unsorted)

	if len(unsorted) < 2 {
		return unsorted
	}


	mid := len(unsorted) / 2;

	 return sort_in_place(merge_in_place(unsorted[0:mid]), merge_in_place(unsorted[mid:]))

}

func sort_in_place (left []int, right []int) []int {
	
	l,r := 0,0;

	for ; l < len(left) && r < len(right); {

		fmt.Println(left[l], right[r])

		l++;
		r++

	}

	res := []int {

	}


	return res

}


func main() {

	var unsorted_arr []int = []int{
		4,3,2,1,5,
	}

	

	// merge(unsorted_arr) // return new array and shorting from existing slice 

	merge_in_place(unsorted_arr)



}
