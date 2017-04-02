package main

import (
	"fmt"
	"os"
	"strconv"
)

func digit(n int) []int {
	digitList := []int{}
	var i int
	for ; n > 0; n = n / 10 {
		i = n % 10
		digitList = append(digitList, i)
		fmt.Println(i, n, digitList)
	}
	return digitList
}

func main() {
	if len(os.Args) > 1 {
		s := os.Args[1]
		if i, err := strconv.Atoi(s); err == nil {
			fmt.Printf("i=%d, type: %T\n", i, i)
			x := digit(i)
			fmt.Println(x)
		}
	}
}
