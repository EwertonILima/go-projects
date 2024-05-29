package main

import (
	"fmt"
	"time"
)


func main() {
	timeStart := time.Now()
	fmt.Println(findSum(9999999999))
	timeEnd := time.Now()

	fmt.Println(timeEnd.Sub(timeStart).Milliseconds())

}


func findSumIterate(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}

func findSum(n int) int {
	return n * (n + 1) / 2
}