package main

import "fmt"

var X = []int{1, 99, 56, 85, 46, 52, 3, 2, 45, 454, 4646, 858, 5, 45, 74, 65, 45}

func main() {
	n := len(X)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if X[min] > X[j] {
				min = j
			}
		}
		X[i], X[min] = X[min], X[i]
	}
	fmt.Print(X)
}
