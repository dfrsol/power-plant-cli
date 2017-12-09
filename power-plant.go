package main

type PlantType string
type PlantStatus string

type PowerPlant struct {
	capacity  float64
	plantType PlantType
	status    PlantStatus
}