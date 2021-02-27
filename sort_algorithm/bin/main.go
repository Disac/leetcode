package main

import (
	"fmt"
	"github.com/disac/leetcode/sort_algorithm"
)

func main() {
	var s = []int{3, 2, 5, 6, 1, 7, 9, 0}
	//fmt.Println(sort_algorithm.BubbleSort(s))
	//fmt.Println(sort_algorithm.InsertSort(s))
	fmt.Println(sort_algorithm.QuickSort(s))
}
