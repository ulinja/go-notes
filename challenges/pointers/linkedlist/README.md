# Code Challenge (Pointers): Linked List Implementation

**Difficulty:** Hard

Implement a simple singly linked list in Go using structs and pointers.
Your implementation should support the following operations:

1. Adding a new node at the end.
2. Deleting a node with a specific value.
3. Printing the list.

## Example
```go
func main() {
    list := &LinkedList{}
    list.Add(10)
    list.Add(20)
    list.Add(30)
    fmt.Println("Original List:")
    list.Print()

    list.Delete(20)
    fmt.Println("After Deleting 20:")
    list.Print()
}
```

**Expected Output:**
```
Original List:
10 -> 20 -> 30 -> nil
After Deleting 20:
10 -> 30 -> nil
```
