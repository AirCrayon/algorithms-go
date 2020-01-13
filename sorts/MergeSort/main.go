package main

import (
	"fmt"
)

func sort(source []int) []int {
	length := len(source)
	if length < 2 {
		return source
	}
	middle := length / 2
	left := source[0:middle]
	right := source[middle:]
	fmt.Println(left, right)
	return merge(sort(left), sort(right))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

func main() {
	source := []int{10, 29, 14, 12, 53, 2, 5}
	result := sort(source)
	fmt.Println(result)
}
