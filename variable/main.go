package main

import "fmt"

func main() {
	var age int // 0
	fmt.Println("My age is", age)
	age = 30
	fmt.Println("My age assignment", age)
	var age1 = 29 // type will be inferred
	fmt.Println("My initial age is", age1)
	var price, quantity int = 5000, 100 //declaring multiple variables
	fmt.Println("price is", price, "quantity is", quantity)

	var (
		name   = "Naveen"
		age2   = 38
		height int
	)
	fmt.Println("my name is", name)
	fmt.Println("my age is", age2)
	fmt.Println("my height is", height)

	count := 10
	fmt.Println("Count =", count)

	name2, age3 := "Naveen", 29 //short hand declaration

	fmt.Println("my name is", name2)
	fmt.Println("my age is", age3)

	a, b := 20, 30 //a and b declared
	fmt.Println("a is", a, "b is", b)
	// a, b := 40, 50 //error, no new variables
}
