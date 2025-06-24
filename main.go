package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	numCh := make(chan []int)
	sqrCh := make(chan []int)
	go generateNum(numCh)
	go squareNum(numCh, sqrCh)
	result := <-sqrCh
	fmt.Println(result)
}

func generateNum(out chan []int) {
	var randSlice []int
	for i := 0; i < 10; i++ {
		num := randNum()
		randSlice = append(randSlice, num)
	}
	out <- randSlice
}

func squareNum(in chan []int, out chan []int) {
	var sqrSlice []int
	for _, val := range <-in {
		sqrNum := int(math.Pow(float64(val), 2))
		sqrSlice = append(sqrSlice, sqrNum)
	}
	out <- sqrSlice
}

func randNum() int {
	return rand.Intn(101)
}
