package main

import (
	"fmt"
)

var plants []PowerPlant
var grid Grid

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)


func init() {
	plants = []PowerPlant{
		{300, hydro, active},
		{45, wind, active},
		{15, wind, inactive},
		{45, solar, unavailable},
		{40, solar, active},
	}

	grid = Grid{300, plants}
}

func main() {
	var option string

	fmt.Println("Choose an option:")
	fmt.Println("1) Generate a Power Plant Report")
	fmt.Println("2) Generate a Grid Report")

	fmt.Scanln(&option)

	switch option {
	case "1":
		grid.GeneratePowerPlantReport()
	case "2":
		grid.GeneratePowerGridReport()
	default:
		fmt.Println("Sorry, you've entered an invalid option")
	}
}
