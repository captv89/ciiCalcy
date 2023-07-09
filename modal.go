package main

// Input Struct for receiving data
type Input struct {
	ImoNumber       string        `json:"imo_number"`
	ShipName        string        `json:"ship_name"`
	ShipType        string        `json:"ship_type"`
	Flag            string        `json:"flag"`
	YearBuilt       string        `json:"year_built"`
	HomePort        string        `json:"home_port"`
	DeadWeight      string        `json:"dead_weight"`
	GrossTonnage    string        `json:"gross_tonnage"`
	FuelConsumption []Consumption `json:"fuel_consumption"`
	RatingYear      string        `json:"rating_year"`
}

type Consumption struct {
	Distance string `json:"distance"`
	Fuel     string `json:"fuel"`
	Quantity string `json:"quantity"`
}

type CIIReference struct {
	ShipTypeCode string `json:"ship_type_code"`
	Case         string `json:"case"`
	Capacity     string `json:"capacity"`
	AValue       string `json:"a_value"`
	CValue       string `json:"c_value"`
}

// Output Struct for sending data
type Output struct {
	ImoNumber        string `json:"imo_number"`
	ShipName         string `json:"ship_name"`
	ShipType         string `json:"ship_type"`
	DeadWeight       string `json:"dead_weight"`
	GrossTonnage     string `json:"gross_tonnage"`
	TotalDistance    string `json:"total_distance"`
	TotalCO2Emission string `json:"total_co2_emission"`
	AttainedCII      string `json:"attained_cii"`
	CIIReference     string `json:"cii_reference"`
	RequiredCII      string `json:"required_cii"`
	CIIScore         string `json:"cii_score"`
	CIIRating        string `json:"cii_rating"`
	RatingYear       string `json:"rating_year"`
}
