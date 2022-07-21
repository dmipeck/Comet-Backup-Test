package main

import "fmt"

type Intersection struct {
	north float32
	east  float32
	south float32
	west  float32
}

func (intersection *Intersection) allRoads() []float32 {
	return []float32{intersection.north, intersection.east, intersection.south, intersection.west}
}

func (intersection *Intersection) cpsTotal() float32 {
	var total float32 = 0
	for _, roadCpm := range intersection.allRoads() {
		total += roadCpm
	}
	return total
}

func main() {
	intersection := Intersection{
		north: 5,
		east:  5,
		south: 5,
		west:  5,
	}
	fmt.Println(intersection.cpsTotal())
}
