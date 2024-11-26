package main

import "fmt"

func reverse(s []int) {
	left, right := 0, len(s)-1

	for {
		s[left], s[right] = s[right], s[left]
		left++
		right--
		if left >= right {
			break
		}
	}

	//for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
	//	s[l], s[r] = s[r], s[l]
	//}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(s)

	reverse(s)

	fmt.Println(s)
}
