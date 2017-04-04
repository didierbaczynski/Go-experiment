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
	}
	return digitList
}

func digPow(n, p int) int {
	x := digit(n)
	fmt.Println(x)
	return 1
}

func main() {
	if len(os.Args) > 2 {
		s1 := os.Args[1]
		s2 := os.Args[2]
		n, err1 := strconv.Atoi(s1)
		p, err2 := strconv.Atoi(s2)
		if err1 == nil && err2 == nil {
			fmt.Printf("n=%d, type: %T n=%d, type: %T\n", n, n, p, p)
			result := digPow(n, p)
			fmt.Printf("Result = %d\n", result)
		}
	}
}
