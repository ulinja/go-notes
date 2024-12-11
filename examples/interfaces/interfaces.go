package main

import (
	"fmt"
)

type SelfIdentifier interface {
	IdentifySelf() string
}

type Human struct {
	Name string
}

func (h *Human) IdentifySelf() string {
	return fmt.Sprintf("Hello, I am %s.", h.Name)
}

type Robot struct {
	SerialNumber uint
}

func (r *Robot) IdentifySelf() string {
	identifier := fmt.Sprintf("0X010-%d", r.SerialNumber)
	return fmt.Sprintf("Beep Boop. I am %s.", identifier)
}

func main() {
	h := Human{"John Smith"}
	r := Robot{4269}

	entities := []SelfIdentifier{&h, &r}
	for _, e := range entities {
		fmt.Println(e.IdentifySelf())
	}
}
