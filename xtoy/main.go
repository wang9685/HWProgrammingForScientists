package main

import (
	"fmt"
	"math/rand"
)

func main() {
	list := RandomGenerate(10)
	fmt.Println(list)
}

func RandomGenerate(n int) []int {
	numList := make([]int, n)
	for i := 0; i < n; i++ {
		num := rand.Intn(2)
		numList = append(numList, num)
	}
	return numList
}