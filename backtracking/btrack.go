package main

import "fmt"

func combine(n int, k int) [][]int {
    result := [][]int{}
    var bTracking func(start int , sliceE[]int)
    bTracking = func(start int , sliceE[]int) {
        if len(sliceE) == k {
            newSlice := make([]int, len(sliceE))
            copy(newSlice, sliceE)
            result = append(result, newSlice)
            return
        }

        for i := start ; i <= n ; i++ {
            sliceE = append(sliceE, i)
            bTracking(i + 1, sliceE)
            sliceE = sliceE[:len(sliceE) -1]
        }

    }

    bTracking(1, []int{})

    return result
    
}

func main () {
	result := combine(5, 1);
	fmt.Println(result)
}