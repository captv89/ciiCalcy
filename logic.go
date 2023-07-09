package main

import (
	"log"
	"math"
	"strconv"
)

// CIICalculator is the function to calculate the CII
func CIICalculator(shipType, deadWeight, grossTonnage, totalDistance, totalCO2Emission, ratingYear string) (attainedCII, ciiReference, requiredCII, ciiScore, ciiRating string) {

	// Convert strings to float64
	totalDistanceFloat, _ := strconv.ParseFloat(totalDistance, 64)
	totalCO2EmissionFloat, _ := strconv.ParseFloat(totalCO2Emission, 64)
	deadWeightFloat, _ := strconv.ParseFloat(deadWeight, 64)
	grossTonnageFloat, _ := strconv.ParseFloat(grossTonnage, 64)

	// Get Ship Type Code
	shipTypeCode := vesselType[shipType]

	// Get the CII reference value
	ciiReference = getCIIReference(shipTypeCode, deadWeight, grossTonnage)
	ciiReferenceFloat, _ := strconv.ParseFloat(ciiReference, 64)

	// Calculate the attained CII
	var attainedCIIFloat float64
	if shipTypeCode < "9" {
		attainedCIIFloat = (math.Pow(10, 6) * totalCO2EmissionFloat) / (deadWeightFloat * totalDistanceFloat)
		attainedCII = strconv.FormatFloat(attainedCIIFloat, 'f', 2, 64)
	} else {
		attainedCIIFloat = (math.Pow(10, 6) * totalCO2EmissionFloat) / (grossTonnageFloat * totalDistanceFloat)
		attainedCII = strconv.FormatFloat(attainedCIIFloat, 'f', 2, 64)
	}

	// Calculate the required CII
	redFactor := reductionFactor[ratingYear]
	redFactorFloat, _ := strconv.ParseFloat(redFactor, 64)
	requiredCIIFloat := ciiReferenceFloat * (100 - redFactorFloat) / 100
	requiredCII = strconv.FormatFloat(requiredCIIFloat, 'f', 2, 64)

	// Calculate the CII score
	ciiScoreFloat := attainedCIIFloat / requiredCIIFloat
	ciiScore = strconv.FormatFloat(ciiScoreFloat, 'f', 2, 64)

	// Calculate the CII rating
	d1 := calcDValues[shipTypeCode][0]
	d2 := calcDValues[shipTypeCode][1]
	d3 := calcDValues[shipTypeCode][2]
	d4 := calcDValues[shipTypeCode][3]

	if ciiScoreFloat < d1 {
		ciiRating = "A"
	} else if ciiScoreFloat < d2 {
		ciiRating = "B"
	} else if ciiScoreFloat < d3 {
		ciiRating = "C"
	} else if ciiScoreFloat < d4 {
		ciiRating = "D"
	} else if ciiScoreFloat > d4 {
		ciiRating = "E"
	} else {
		log.Fatal("Error in calculating CII rating")
	}

	return
}

// getCIIReference is the function to get the CII reference value
func getCIIReference(shipTypeCode, deadWeight, grossTonnage string) (ciiReference string) {

	var ciiRef float64

	// Convert strings to float64
	deadWeightFloat, _ := strconv.ParseFloat(deadWeight, 64)
	grossTonnageFloat, _ := strconv.ParseFloat(grossTonnage, 64)

	// Calculate the CII reference value
	if shipTypeCode == "1" && deadWeightFloat >= 279000 {
		aValue := 4745
		cValue := -0.622
		ciiRef = float64(aValue) * math.Pow(279000, cValue)

	} else if shipTypeCode == "1" && deadWeightFloat < 279000 {
		aValue := 4745
		cValue := -0.622
		ciiRef = float64(aValue) * math.Pow(deadWeightFloat, cValue)
	} else if shipTypeCode == "2" && deadWeightFloat >= 65000 {
		aValue := 144050000000
		cValue := -2.071
		ciiRef = float64(aValue) * math.Pow(deadWeightFloat, cValue)
	} else if shipTypeCode == "2" && deadWeightFloat < 65000 {
		aValue := 8104
		cValue := -0.639
		ciiRef = float64(aValue) * math.Pow(deadWeightFloat, cValue)
	} else if shipTypeCode == "3" {
		aValue := 5247
		cValue := -0.610
		ciiRef = float64(aValue) * math.Pow(grossTonnageFloat, cValue)
	} else if shipTypeCode == "4" {
		aValue := 1984
		cValue := -0.489
		ciiRef = float64(aValue) * math.Pow(deadWeightFloat, cValue)
	} else if shipTypeCode == "5" && deadWeightFloat >= 20000 {
		aValue := 31948
		cValue := -0.792
		ciiRef = float64(aValue) * math.Pow(deadWeightFloat, cValue)
	} else if shipTypeCode == "5" && deadWeightFloat < 20000 {
		aValue := 588
		cValue := -0.3885
		ciiRef = float64(aValue) * math.Pow(deadWeightFloat, cValue)
	} else if shipTypeCode == "6" {
		aValue := 4600
		cValue := -0.557
		ciiRef = float64(aValue) * math.Pow(deadWeightFloat, cValue)
	} else if shipTypeCode == "7" {
		aValue := 5119
		cValue := -0.622
		ciiRef = float64(aValue) * math.Pow(deadWeightFloat, cValue)
	} else if shipTypeCode == "8" && deadWeightFloat >= 100000 {
		aValue := 9.827
		cValue := 0
		ciiRef = aValue * math.Pow(deadWeightFloat, float64(cValue))
	} else if shipTypeCode == "8" && deadWeightFloat >= 65000 {
		aValue := 144790000000000
		cValue := -2.673
		ciiRef = float64(aValue) * math.Pow(deadWeightFloat, cValue)
	} else if shipTypeCode == "8" && deadWeightFloat < 65000 {
		aValue := 144790000000000
		cValue := -2.673
		ciiRef = float64(aValue) * math.Pow(65000, cValue)
	} else if shipTypeCode == "9" && grossTonnageFloat >= 57700 {
		aValue := 3672
		cValue := -0.590
		ciiRef = float64(aValue) * math.Pow(57700, cValue)
	} else if shipTypeCode == "9" && grossTonnageFloat >= 30000 {
		aValue := 3672
		cValue := -0.590
		ciiRef = float64(aValue) * math.Pow(grossTonnageFloat, cValue)
	} else if shipTypeCode == "9" && grossTonnageFloat < 30000 {
		aValue := 330
		cValue := -0.329
		ciiRef = float64(aValue) * math.Pow(grossTonnageFloat, cValue)
	} else if shipTypeCode == "10" {
		aValue := 1967
		cValue := -0.485
		ciiRef = float64(aValue) * math.Pow(grossTonnageFloat, cValue)
	} else if shipTypeCode == "11" {
		aValue := 2023
		cValue := -0.460
		ciiRef = float64(aValue) * math.Pow(grossTonnageFloat, cValue)
	} else if shipTypeCode == "12" {
		aValue := 4196
		cValue := -0.460
		ciiRef = float64(aValue) * math.Pow(grossTonnageFloat, cValue)
	} else if shipTypeCode == "13" {
		aValue := 930
		cValue := -0.383
		ciiRef = float64(aValue) * math.Pow(grossTonnageFloat, cValue)
	} else {
		ciiRef = 0
		log.Println("No ship type found")
	}

	ciiReference = strconv.FormatFloat(ciiRef, 'f', 2, 64)
	return ciiReference
}

// TotalConsumption is the function to calculate the total consumption
func TotalConsumption(fuelConsumption []Consumption) (totalDistance, totalCO2Emission string) {
	totalDist := 0.0
	totalCO2 := 0.0
	// Loop through the fuel consumption
	for _, consumption := range fuelConsumption {

		// Add the distance
		distance, _ := strconv.ParseFloat(consumption.Distance, 64)
		totalDist += distance

		// Calculate the CO2 emission
		// Convert strings to float64
		quantity, _ := strconv.ParseFloat(consumption.Quantity, 64)
		// Get the factor value from mapping
		factor := fuelType[consumption.Fuel]
		// Convert strings to float64
		factorFloat, _ := strconv.ParseFloat(factor, 64)
		totalCO2 += quantity * factorFloat
	}

	// Convert float64 to string
	totalDistance = strconv.FormatFloat(totalDist, 'f', 2, 64)
	totalCO2Emission = strconv.FormatFloat(totalCO2, 'f', 2, 64)

	return totalDistance, totalCO2Emission
}
