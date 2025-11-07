package main

import "fmt"

func main() {

	var arr []int = []int{
		1,2,3,4,5,
	}


	fmt.Println(binary_search(arr, 5, 0, len(arr) - 1))

}

func binary_search(arr []int, target int, low int, high int ) int {

	if len(arr) < 1 || low > high {
		return -1
	}

	mid := (low + high) / 2;
	if mid < 0 || mid >= len(arr) {
		return -1
	}

	if arr[mid] == target {
		return mid
	}

	if arr[mid] > target {
		return binary_search(arr, target, low, mid - 1)
	}

	
	return binary_search(arr, target, mid + 1 , high)

}
