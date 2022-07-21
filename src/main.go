package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

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
	Name           string  `json:"name"`
	EfficiencyHigh float32 `json:"efficiency_high"`
	EfficiencyMed  float32 `json:"efficiency_med"`
	EfficiencyLow  float32 `json:"efficiency_low"`
}

func (control *ControlMethod) cpsEfficency(intersection Intersection) float32 {
	cpsTotal := intersection.cpsTotal()
	if cpsTotal < 10 {
		return control.EfficiencyLow
	} else if cpsTotal < 20 {
		return control.EfficiencyMed
	} else {
		return control.EfficiencyHigh
	}
}

func main() {
	northInPtr := flag.Float64("north", 0, "north road flow in CPM")
	eastInPtr := flag.Float64("east", 0, "east road flow in CPM")
	southInPtr := flag.Float64("south", 0, "south road flow in CPM")
	westInPtr := flag.Float64("west", 0, "west road flow in CPM")
	flag.Parse()

	file, err := ioutil.ReadFile("intersections.json")
	if err != nil {
		fmt.Println("Failed to open \"intersections.json\"")
		panic(err)
	}

	var controlMethods []ControlMethod
	err = json.Unmarshal([]byte(file), &controlMethods)
	if err != nil {
		fmt.Println("Failed to parse json from \"intersections.json\"")
		panic(err)
	}

	intersection := Intersection{
		north: float32(*northInPtr),
		east:  float32(*eastInPtr),
		south: float32(*southInPtr),
		west:  float32(*westInPtr),
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
	fmt.Printf("Intersection Best Method: %s\n", bestMethod.Name)
	fmt.Printf("Intersection Best Method Efficency: %.2f%%\n", bestEfficency)
	fmt.Printf("Intersection Best Method CPM: %.2f\n", intersection.cpsTotal()*bestEfficency)
}
