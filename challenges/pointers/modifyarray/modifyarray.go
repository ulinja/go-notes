package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Array:", arr)
	DoubleArray(&arr)
	fmt.Println("Modified Array:", arr)
}

func DoubleArray(arr *[]int) {
	const l = 5
	a := *arr
	for i := 0; i < l; i++ {
		a[i] *= 2
	}
}
