package main

import "fmt"
import "github.com/mitchellh/mapstructure"

type Point struct {
	X int
	Y int
}

func (p Point) method() {
	fmt.Println(p.X, p.Y)
}

func main() {
	//pointsMap := map[string]Point{}
	//otherPointsMap := make(map[int]Point)
	//oneMorePointsMap := map[string]Point
	//pointsMap["a"] = Point{
	//	X: 1,
	//	Y: 2,
	//}
	//
	//otherPointsMap[1] = Point{X: 1, Y: 2}
	//fmt.Println(pointsMap)
	//fmt.Println(pointsMap["a"])
	//fmt.Println(otherPointsMap)

	//var oneMorePointsMap map[string]Point
	//
	//if oneMorePointsMap == nil {
	//	oneMorePointsMap = map[string]Point{
	//		"a": {X: 1, Y: 2},
	//	}
	//	fmt.Println(oneMorePointsMap)
	//}
	//oneMorePointsMap["b"] = Point{1, 2}
	//oneMorePointsMap["b"].method()
	//
	//value, ok := oneMorePointsMap["b"]
	//if ok {
	//	fmt.Println(value, ok)
	//}

	pointsMap := map[string]int{
		"X": 1,
		"Y": 2,
	}

	p1 := Point{}
	mapstructure.Decode(pointsMap, &p1)
	fmt.Println(p1)
}
