package main

import "fmt"

func maxNonDuplicated(s string) int  {
	var charMap map[byte]int = map[byte]int{}
	left,maxLen := 0,0;

	for right := 0; right < len(s); right++ {
		indexCharCode, isExist := charMap[s[right]];
		if (!isExist) {
			charMap[s[right]] = right;
		} else {
			// fmt.Printf("%v ada di index %d ",s[right], indexCharCode);
			// fmt.Println(indexCharCode, isExist)
			if (indexCharCode >= left) {		
				left = indexCharCode + 1
				charMap[s[right]] = right
			} else {
				charMap[s[right]] = right
			}
		}
		if (right - left + 1) > maxLen {
			maxLen = right - left + 1
		}
		fmt.Printf("%d - %d = %d\n", right , left, right - left + 1)
	}

	// fmt.Println(maxLen)
	

	return maxLen

}

func main() {

	test := "dvdaf"
	fmt.Println(len(test))


	maxNonDuplicated(test)

}