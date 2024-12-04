package main

import (
	"fmt"
)

func main() {
	data := []int{1, 1, 2, 2, 3, 3}

	dist := make(map[[2]int]int)

	dist[([2]int)(data[0:2])] = 10
	dist[([2]int)(data[2:4])] = 20
	dist[([2]int)(data[4:6])] = 30

	fmt.Println(dist)

	points := getPoints(data)

	fmt.Println(points)

}

func getPoints(data []int) [][2]int {
	res := make([][2]int, 0, 3)

	for i := 0; i < len(data); i += 2 {
		res = append(res, ([2]int)(data[i:i+2]))
	}

	return res
}
