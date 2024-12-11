package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Println("Before Swap: a =", a, "b =", b)
	Swap(&a, &b)
	fmt.Println("After Swap: a =", a, "b =", b)
}

func Swap(a, b *int) {
	x := *a
	*a = *b
	*b = x
}
