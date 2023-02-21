package main

import "fmt"

var aa = [...]string{"abcdefg", "1234", "啊"}

func main() {

	for i := len(aa) - 1; i >= 0; i-- {

		fmt.Print(aa[i], " ")
	}

}
