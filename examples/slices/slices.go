package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    s = s[:0]
    printSlice(s)

    s = s[:4]
    printSlice(s)

    s = s[2:]
    printSlice(s)

    nums := []int{2, 4, 6, 8, 10}
    for i, v := range nums {
        fmt.Printf("Value at %d is %d\n", i, v)
    }
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
