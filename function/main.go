package main

import "fmt"

func calculateBill(price, quantity int) int {
	var totalPrice = price * quantity
	return totalPrice
}

func rectProps(length, width float32) (float32, float32) {

	var area = length * width
	var perimeters = (length + width) * 2

	return area, perimeters
}

func rectProps2(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

func rectProps3(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	price, quantity := 90, 6

	total := calculateBill(price, quantity)

	fmt.Println("Total price:", total)

	area, perimeter := rectProps(10.8, 5.6)

	fmt.Printf("Area %f Perimeter %f", area, perimeter)
	fmt.Println()
	area3, _ := rectProps3(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f ", area3)
}
