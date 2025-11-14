package main;
import "fmt"

func canJump(nums []int) bool {
    if len(nums) <= 1 {
        return true
    }
	maxJump := nums[0];
	for i , v := range nums {

		if i < len(nums) - 1 && maxJump >= i {
			if v + i >= maxJump {
				maxJump = v + i
                if maxJump >= len(nums) - 1 {
		        return true
	            } 
			}
		}

	}

	return false
}

func main () {

	fmt.Println(canJump([]int{
		1,2,2,0,
	}))

}

// this solution is linear time complexity O(n) space O(1)