package main

import "fmt"

type Point struct {
	X, Y int
}

func movePoint(p *Point, x, y int) {
	p.X += x
	p.Y += y
}

func movePointPtr(p *Point, x, y int) {
	p.X += x
	p.Y += y
}

func main() {
	p := Point{1, 2}
	//fmt.Println(p.movePoint(1, 2))
	//fmt.Println(movePoint(p, 3, 4))
	//p.movePointPtr(3, 4)
	//fmt.Println(p.X, p.Y)

	v := &p

	movePoint(v, 1, 2)
	fmt.Println(*v)
	movePointPtr(v, 2, 3)
	fmt.Println(*v)
}
