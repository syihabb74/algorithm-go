package main;
import "fmt"

func jump(nums []int) int {
    	r,max,minStep := nums[0],0,0
	if r == 0 || len(nums) < 2 {
		return 0
	}
	for k,v := range nums {

		if k + v >= max {
			max = k + v
		}

		if (k == r || k == len(nums) - 1) {
			minStep++
			r = max
			max = 0
		}
	}
    return minStep
}

// above is big O(n) time complexity (linear time complexity)


func main () {
// call here jump func if you wanna test mine 

}
