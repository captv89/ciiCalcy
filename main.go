package main

import (
	"fmt"
	"log"
)

// Main function
func main() {

	//	Sample input
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

// validateInput is the function to validate the input
func validateInput(input Input) bool {
	shipType := input.ShipType

	// Check if the ship type is valid
	_, ok := vesselType[shipType]
	if !ok {
		log.Println("Invalid ship type")
		return true
	}

	// Check if the fuel consumption is valid
	for _, fuelConsumption := range input.FuelConsumption {
		fuel := fuelConsumption.Fuel
		_, ok := fuelType[fuel]
		if !ok {
			log.Println("Invalid fuel type")
			return true
		}
	}

	// Check if the rating year is valid
	_, ok = reductionFactor[input.RatingYear]
	if !ok {
		log.Println("Invalid rating year")
		return true
	}

	return false
}

// Action is the function to read the input from the user
func Action(input Input) {

	// Call the TotalConsumption function to calculate the total distance and total CO2 emission
	totalDistance, totalCO2Emission := TotalConsumption(input.FuelConsumption)

	// Call the CIICalculator function to calculate the CII
	attainedCII, ciiReference, requiredCII, ciiScore, ciiRating := CIICalculator(input.ShipType, input.DeadWeight, input.GrossTonnage, totalDistance, totalCO2Emission, input.RatingYear)

	// Set the output
	output := Output{
		ImoNumber:        input.ImoNumber,
		ShipName:         input.ShipName,
		ShipType:         input.ShipType,
		DeadWeight:       input.DeadWeight,
		GrossTonnage:     input.GrossTonnage,
		TotalDistance:    totalDistance,
		TotalCO2Emission: totalCO2Emission,
		AttainedCII:      attainedCII,
		CIIReference:     ciiReference,
		RequiredCII:      requiredCII,
		CIIScore:         ciiScore,
		CIIRating:        ciiRating,
		RatingYear:       input.RatingYear,
	}

	// Print the output
	fmt.Println(output)

	return
}
