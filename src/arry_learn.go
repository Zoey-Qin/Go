package main

import (
	"fmt"
)

func main() {
	test := [][]int{}
	a1 := []int{1,2,3,4}
	a2 := []int{5,6,7,8}
	test = append(test, a1)
	test = append(test, a2)
	fmt.Println("get all test data by index:")
	for i := 0; i<2; i++{
		fmt.Println(test[i])
	}
	fmt.Println("get a1[3]")
	fmt.Println(test[0][3])
}
