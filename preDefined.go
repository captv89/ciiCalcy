package main

// Predefine the recognised vessel type and their corresponding codes
var vesselType = map[string]string{
	"Bulk carrier":                            "1",
	"Gas carrier":                             "2",
	"Tanker":                                  "3",
	"Container ship":                          "4",
	"General cargo ship":                      "5",
	"Refrigerated cargo carrier":              "6",
	"Combination carrier":                     "7",
	"LNG carrier":                             "8",
	"Ro-Ro cargo ship (Vehicle Carrier)":      "9",
	"Ro-Ro cargo ship":                        "10",
	"Ro-Ro passenger ship":                    "11",
	"Ro-Ro passenger ship (High speed craft)": "12",
	"Cruise passenger ship":                   "13",
}

// Predefine fuel types and their factors
var fuelType = map[string]string{
	"Diesel/Gas Oil": "3.206",
	"LFO":            "3.151",
	"HFO":            "3.114",
	"LPG(Propane)":   "3.000",
	"LPG(Butane)":    "3.030",
	"LNG":            "2.750",
	"Methanol":       "1.375",
	"Ethanol":        "1.913",
}

// Predefine reduction factor for each year
var reductionFactor = map[string]string{
	"2023": "5",
	"2024": "7",
	"2025": "9",
	"2026": "11",
}

// Predefine the d values
var calcDValues = map[string][]float64{
	"1":  {0.86, 0.94, 1.06, 1.18},
	"2":  {0.81, 0.91, 1.12, 1.44},
	"3":  {0.82, 0.93, 1.08, 1.28},
	"4":  {0.83, 0.94, 1.07, 1.19},
	"5":  {0.83, 0.94, 1.06, 1.19},
	"6":  {0.78, 0.91, 1.07, 1.2},
	"7":  {0.87, 0.96, 1.06, 1.14},
	"8":  {0.89, 0.98, 1.06, 1.13},
	"9":  {0.86, 0.94, 1.06, 1.16},
	"10": {0.76, 0.89, 1.08, 1.27},
	"11": {0.76, 0.92, 1.14, 1.3},
	"12": {0.76, 0.92, 1.14, 1.3},
	"13": {0.87, 0.95, 1.06, 1.16},
}
