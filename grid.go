package main

import (
	"fmt"
	"strings"
)

type Grid struct {
	load   float64
	plants []PowerPlant
}

func (g *Grid) GeneratePowerPlantReport() {
	for i, plant := range g.plants {
		label := fmt.Sprintf("%s%d", "Plant #", i)

		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%s\n", "Type", plant.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity", plant.capacity)
		fmt.Printf("%-20s%s\n", "Status", plant.status)
		fmt.Println("")
	}
}

func (g *Grid) GeneratePowerGridReport() {
	capacity := 0.

	for _, plant := range g.plants {
		if plant.status == active {
			capacity += plant.capacity
		}
	}

	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%.0f\n", "Capacity", capacity)
	fmt.Printf("%-20s%.0f\n", "Current Load", g.load)
	fmt.Printf("%-20s%.2f\n", "Utilization", g.load/capacity*100)
}

