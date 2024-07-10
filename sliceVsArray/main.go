package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}
	printSlice(slice)
	fmt.Println("\n", len(slice))


	var array = [6]int{1, 2, 3}
	printArray(array)
	fmt.Println("\n", cap(array))
}

func printSlice(list []int) {
	for _, item := range list {
		fmt.Printf("%d\t", item)
	}
}

func printArray(list [6]int) {
	for _, item := range list {
		fmt.Printf("%d\t", item)
	}
}
