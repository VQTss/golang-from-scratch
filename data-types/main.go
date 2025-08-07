package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	var f = 89
	g := 95
	fmt.Println("value of f is", f, "and g is", g)
	fmt.Printf("type of f is %T, size of f is %d bytes", f, unsafe.Sizeof(f))   //type and size of a
	fmt.Printf("\ntype of g is %T, size of g is %d bytes", g, unsafe.Sizeof(g)) //type and size of b

	var e uint = 60
	var r uint = 30
	t := e * r
	fmt.Println("t =", t)
	fmt.Printf("Data type of variable t is %T", t)
}
