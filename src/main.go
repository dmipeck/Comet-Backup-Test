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

func printHelp() {
	fmt.Println("Comet Backup Test Program:")
	fmt.Println("calculates the best control method for an intersection")
	fmt.Println("the available control methods can be edited by editing the \"control_methods.json\" file")
	fmt.Println("\tflags:")
	fmt.Println("\t\t-help\t prints help")
	for _, direction := range []string{"north", "east", "south", "west"} {
		fmt.Printf("\t\t-%s=<number>\t sets the %s CPM to the given value\n", direction, direction)
	}
}

func main() {
	printHelpInPtr := flag.Bool("help", false, "Print help")
	northInPtr := flag.Float64("north", 0, "north road flow in CPM")
	eastInPtr := flag.Float64("east", 0, "east road flow in CPM")
	southInPtr := flag.Float64("south", 0, "south road flow in CPM")
	westInPtr := flag.Float64("west", 0, "west road flow in CPM")
	flag.Parse()

	if *printHelpInPtr {
		printHelp()
		return
	}

	fileName := "control_methods.json"
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Failed to open \"%s\"\n", fileName)
		panic(err)
	}

	var controlMethods []ControlMethod
	err = json.Unmarshal([]byte(file), &controlMethods)
	if err != nil {
		fmt.Printf("Failed to parse json from \"%s\"\n", fileName)
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
