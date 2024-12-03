# Code Challenge (Pointers): Number Swap

**Difficulty:** Easy

Write a function `Swap` that takes two integer pointers as arguments and swaps the values they point to.
The main function should test this with at least three different sets of numbers.

## Example
```go
func main() {
    a, b := 5, 10
    fmt.Println("Before Swap: a =", a, ", b =", b)
    Swap(&a, &b)
    fmt.Println("After Swap: a =", a, ", b =", b)
}
```

**Expected Output:**
```
Before Swap: a = 5, b = 10
After Swap: a = 10, b = 5
```
