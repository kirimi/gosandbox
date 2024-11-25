package main

import "fmt"

func foo(src []int) {
	src = append(src, 5)
}

func main() {
	arr := []int{1, 2, 3}
	src := make([]int, 1)
	copy(src, arr)

	foo(src)

	printSlice(arr)
	printSlice(src)

}

func printSlice(slice []int) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	fmt.Printf("%v\n", slice)
}
