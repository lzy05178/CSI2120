package main

import (
	"fmt"
	"math"
)

var cha chan int

//var count int = 0

type Point struct {
	x float64
	y float64
}

func MidPoint(point1 Point, point2 Point) {
	fmt.Printf("Points: (%v,%v) (%v,%v)\n", point1.x, point1.y, point2.x, point2.y)

	midpoint := Point{(point1.x + point2.x) / 2, (point1.y + point2.y) / 2}
	fmt.Printf("MidPoint: (%.1f,%.1f)\n", midpoint.x, midpoint.y)

	length := math.Sqrt(math.Pow(point1.x-point2.x, 2) + math.Pow(point1.y-point2.y, 2))
	s := fmt.Sprintf("Length: %.2f", length)
	fmt.Println(s)
	fmt.Println(" ")
	cha <- 0

}

func main() {

	//MidPoint(Point{2, 4}, Point{6, 8})
	//println("")
	//println("***************************")
	//println("")
	//
	//MidPoint(Point{13, 2}, Point{7, 10})
	//println("")
	//println("***************************")
	//println("")
	//
	//MidPoint(Point{3, 3}, Point{-4, -7})

	cha = make(chan int)
	points := []Point{{8., 1.}, {3., 2.}, {7., 4.}, {6., 3.}}
	fmt.Printf("point= %v\n", points[2])

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			go MidPoint(points[i], points[j])
			//count++
		}
	}
	var i = 0
	for i < 6 {
		<-cha
		i++
	}
	//fmt.Println(count)

}
