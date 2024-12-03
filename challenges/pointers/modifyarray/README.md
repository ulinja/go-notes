# Code Challenge (Pointers): Modify Array In-Place

**Difficulty:** Medium

Write a program that creates an array of integers dynamically, modifies its values using pointers, and prints the results.
Use a function `DoubleArray` that takes a pointer to an array of integers and doubles each element's value.

## Example
```go
func main() {
    arr := []int{1, 2, 3, 4, 5}
    fmt.Println("Original Array:", arr)
    DoubleArray(&arr)
    fmt.Println("Modified Array:", arr)
}
```

**Expected Output:**
```
Original Array: [1 2 3 4 5]
Modified Array: [2 4 6 8 10]
```
