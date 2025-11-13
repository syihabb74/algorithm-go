package main;

func twoSum(numbers []int, target int) []int {
    r := len(numbers) - 1;
	for l := 0; l < len(numbers) ;  {
		if (l > r) {
			return []int{

			}
		}
		if numbers[l] + numbers[r] == target {
			return []int{
				l+1, r+1,
			}
		} else if numbers[l] + numbers[r] > target {
			r--
		} else {
			l++
		}

	}
	return []int{}
}

func main () {
	// call two sum here :)
}