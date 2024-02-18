package main

import (
	"fmt"
)

// У структуры только поля
type Point struct {
	X int
	Y int
	S string
}

// так в структуру пережаются методы

func (p Point) getAllPoints() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)
}

func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
		S: "Pubma",
	}

	p1.getAllPoints()
}

// ... - посчитать кол-во элементов в массиве
func workWithArray() {
	var a [2]string
	a[0] = "hello"
	a[1] = "pubma"

	fmt.Println(a)

	b := [...]int{1, 2, 3}
	fmt.Println(b)
}

func workWithSlice() {
	a := []string{"pumba", "kak", "dela"}
	fmt.Println(a)

	b := make([]int, 3)

	for i := 0; i < len(b); i++ {
		b[i] = i
	}
	//b[0] = "a"
	//b[1] = "b"
	//b[2] = "c"
	//b = append(b, "a", "b", "c", "d")
	fmt.Println(b)

	//animalsArray := [4]string{
	//	"dog",
	//	"cat",
	//	"cow",
	//	"elephant",
	//}

	animalsSlice := []string{
		"dog",
		"cat",
		"cow",
	}

	ab := animalsSlice[0:2]
	bc := animalsSlice[1:3]
	fmt.Println(ab)
	fmt.Println(bc)

	t := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	tt := t[5:]
	ttt := t[:7]
	//tttt := t[:]

	fmt.Println(tt)
	fmt.Println(ttt)
	//fmt.Println(tttt)

}

func callbackFunction(callback func(int, int) int, s string, num1 int, num2 int) int {
	result := callback(num1, num2)

	fmt.Println(s)
	return result
}

func main() {
	structs()
	workWithArray()
	workWithSlice()
	sumCallback := func(n1, n2 int) int {
		return n1 + n2
	}

	multipleCallback := func(n1, n2 int) int {
		return n1 * n2
	}

	result := callbackFunction(sumCallback, "plus", 1, 3)
	result2 := callbackFunction(multipleCallback, "multiple", 2, 3)

	fmt.Println(result)
	fmt.Println(result2)

	sales := []int{
		25,
		30,
		40,
		50,
	}

	calculatePrice := totalPrice(100)

	for i := 0; i < len(sales); i++ {
		if sales[i] > 25 {
			fmt.Println(calculatePrice(sales[i]))
		}
	}

}

//Замыкание

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}
