package main

import "fmt"

func gcd(num1 int, num2 int) int {
	if num2 == 0 {
		return num1
	} else {
		return gcd(num2, num1%num2)
	}
}

func countItem(apples int, cakes int, divisor int) (int, int) {
	return apples / divisor, cakes / divisor
}

func main() {
	var apples int
	var cakes int

	fmt.Print("apples = ")
	fmt.Scanln(&apples)

	fmt.Print("cakes  = ")
	fmt.Scanln(&cakes)

	var boxes = gcd(apples, cakes)
	var applesEach, cakesEach = countItem(apples, cakes, boxes)

	fmt.Println("Ainun got", boxes, "boxes")
	fmt.Println("Each box has", applesEach, "apples and", cakesEach, "cakes")

}
