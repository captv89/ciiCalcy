package main

import "fmt"

// Main function
func main() {

	//	Sample input
	input := Input{
		ImoNumber:    "1234567",
		ShipName:     "Sample Ship",
		ShipType:     "Bulk carrier",
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

	Action(input)
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
