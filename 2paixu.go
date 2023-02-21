package main

import "fmt"

var X = []int{1, 99, 56, 85, 46, 52, 3, 2, 45, 454, 4646, 858, 5, 45, 74, 65, 45}

func main() {
	n := len(X)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if X[i] > X[j] {
				X[i], X[j] = X[j], X[i]
			}

		}
	}
	fmt.Print(X)
}
