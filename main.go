package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	numCh := make(chan int)
	sqrCh := make(chan int)
	go generateNum(numCh)
	go squareNum(numCh, sqrCh)
	var result []int
	for val := range sqrCh {
		result = append(result, val)
	}
	fmt.Println(result)
}

func generateNum(out chan int) {
	var randSlice []int
	for i := 0; i < 10; i++ {
		num := randNum()
		randSlice = append(randSlice, num)
	}
	for _, val := range randSlice {
		out <- val
	}
	close(out)
}

func squareNum(in chan int, out chan int) {
	for val := range in {
		sqrNum := int(math.Pow(float64(val), 2))
		out <- sqrNum
	}
	close(out)
}

func randNum() int {
	return rand.Intn(101)
}
