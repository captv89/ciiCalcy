package main

import (
	"log"
	"testing"
)

// Add the testing for the main.go file here
func TestMain(m *testing.M) {
	// Create a sample input
	input := Input{
		ImoNumber:    "1234567",
		ShipName:     "Box Ship",
		ShipType:     "Container ship",
		Flag:         "Singapore",
		YearBuilt:    "2010",
		HomePort:     "Singapore",
		DeadWeight:   "300000",
		GrossTonnage: "280000",
		FuelConsumption: []Consumption{
			{
				Distance: "19000",
				Fuel:     "HFO",
				Quantity: "2800",
			},
			{
				Distance: "1000",
				Fuel:     "Diesel/Gas Oil",
				Quantity: "500",
			},
		},
		RatingYear: "2023",
	}

	// Check if the input is valid
	skipAction := validateInput(input)

	if !skipAction {
		Action(input)
	} else {
		log.Println("Invalid input")
	}
}
