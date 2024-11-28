package main

import (
	"fmt"
	"slices"
	"sort"
)

func intersectWithMap(s1 []int, s2 []int) []int {
	result := make([]int, 0)

	counter := make(map[int]int)

	for _, val := range s1 {
		_, exist := counter[val]
		if exist {
			counter[val]++
		} else {
			counter[val] = 1
		}
	}

	for _, val := range s2 {
		c, exist := counter[val]
		if exist && c > 0 {
			result = append(result, val)
			counter[val]--
		}
	}

	return result
}

func intersectWithSort(s1 []int, s2 []int) []int {
	result := make([]int, 0)
	slices.Sort(s1)
	slices.Sort(s2)

	maxIndex := min(len(s1), len(s2))
	p1, p2 := 0, 0
	for {
		if s1[p1] == s2[p2] {
			result = append(result, s1[p1])
			p1++
			p2++

		} else if s1[p1] < s2[p2] {
			p1++
		} else {
			p2++
		}

		if p1 > maxIndex || p2 > maxIndex {
			break
		}
	}

	return result
}

func main() {
	slice1 := []int{1, 1, 2, 1, 4, 3, 5, 6, 6, 6, 6, 6}
	slice2 := []int{1, 5, 6, 3, 9, 10, 6, 6, 5, 1}

	intersection1 := intersectWithMap(slice1, slice2)

	intersection2 := intersectWithSort(slice1, slice2)

	sort.Ints(intersection1)
	sort.Ints(intersection2)

	isEqual := slices.Equal(intersection1, intersection2)

	fmt.Printf("intersectWithMap: %v\n"+
		"intersectWithSort: %v\n"+
		"isEqual = %t\n",
		intersection1,
		intersection2,
		isEqual,
	)
}
