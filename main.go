package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

// Main function
func main() {
	// Start a gin server
	go startGinServer()

	// Listen to the input to exit the program if exit is typed
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Println(err)
		return
	}

	if input == "exit" {
		log.Println("Exiting the program")
		os.Exit(0)
	}
}

// Start a gin server to listen on 8080
func startGinServer() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default())

	// Handle the POST request
	router.POST("/calculate-cii", calcy)

	// Start the server
	//Start and run the server if production environment
	if os.Getenv("APP_ENV") == "prod" {
		log.Println("Starting server in production environment")
		err := router.RunTLS(fmt.Sprintf(":%s", os.Getenv("PORT")), os.Getenv("CERT_PATH"), os.Getenv("KEY_PATH"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Starting server in development environment")
		err := router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Handle the POST request
func calcy(c *gin.Context) {

	// Process the input
	var input Input
	err := c.ShouldBindJSON(&input)
	// Check for any errors during input processing
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//// Print the input
	log.Println("Input: ", input)

	// Check if the input is valid
	skipAction := validateInput(input)

	if !skipAction {
		output := Action(input)
		c.JSON(http.StatusOK, output)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
	}
}

// validateInput is the function to validate the input
func validateInput(input Input) bool {
	shipType := input.ShipType
	//log.Println("Ship type: ", shipType)
	// Check if the ship type is valid
	_, ok := vesselType[shipType]
	if !ok {
		log.Println("Invalid ship type")
		return true
	}

	// Check if the fuel consumption is valid
	for _, fuelConsumption := range input.FuelConsumptions {
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
func Action(input Input) Output {

	// Call the TotalConsumption function to calculate the total distance and total CO2 emission
	totalDistance, totalCO2Emission := TotalConsumption(input.FuelConsumptions)

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

	return output
}
