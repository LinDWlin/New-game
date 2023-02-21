package main

import "fmt"

var arr = []int{1, 99, 56, 85, 46, 52, 3, 2, 45, 454, 4646, 858, 5, 45, 74, 65, 45}

func main() {

	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
	fmt.Print(arr)
}
