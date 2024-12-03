package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

var (
    v1 = Vertex{1, 2}
    v2 = Vertex{X: 1}
    v3 = Vertex{}
    p = &Vertex{1, 2}
)

func main() {
    v := Vertex{1, 2}
    fmt.Println(v.X)
    fmt.Println(v1, v2, v3, p)
}
