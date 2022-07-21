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

type ControlMethod struct {
	name           string
	efficiencyHigh float32
	efficiencyMed  float32
	efficiencyLow  float32
}

func (control *ControlMethod) cpsEfficency(intersection Intersection) float32 {
	cpsTotal := intersection.cpsTotal()
	if cpsTotal < 10 {
		return control.efficiencyLow
	} else if cpsTotal < 20 {
		return control.efficiencyMed
	} else {
		return control.efficiencyHigh
	}
}

var controlMethods []ControlMethod = []ControlMethod{
	{
		name:           "Roundabout",
		efficiencyHigh: 0.50,
		efficiencyMed:  0.75,
		efficiencyLow:  0.09,
	},
	{
		name:           "Stop Signs",
		efficiencyHigh: 0.20,
		efficiencyMed:  0.30,
		efficiencyLow:  0.40,
	},
	{
		name:           "Traffic Lights",
		efficiencyHigh: 0.90,
		efficiencyMed:  0.75,
		efficiencyLow:  0.30,
	},
}

func main() {
	intersection := Intersection{
		north: 5,
		east:  5,
		south: 5,
		west:  5,
	}

	var bestMethod *ControlMethod
	var bestEfficency float32
	var nextMethodEfficency float32
	for _, nextMethod := range controlMethods {
		nextMethodEfficency = nextMethod.cpsEfficency(intersection)
		if bestMethod == nil || nextMethodEfficency < bestEfficency {
			bestEfficency = nextMethodEfficency
			bestMethod = &nextMethod
		}
	}

	fmt.Printf("Intersection CPM Total: %.2f\n", intersection.cpsTotal())
	fmt.Printf("Intersection Best Method: %s\n", bestMethod.name)
	fmt.Printf("Intersection Best Method Efficency: %.2f%%\n", bestMethod.cpsEfficency(intersection))
}
