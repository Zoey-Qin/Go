package main

import (
	"fmt"
	"log"
)

func main() {
	s1 := make([]int, 3)
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	s1 = append(s1, 4)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	log.Info("test")
	for i := 0; i<len(s1); i++{
		fmt.Println(s1[i])
	}
}
