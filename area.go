package converter

const (
	// Kmmi is a value used to convert Kilometers to Miles
	Kmmi = 0.62137
	// Yame is a value used to convert Yards and Miles
	Yame = 1.0936
	// Ftya is used to convert feet to yards
	Ftya = 0.33333
)

// MilestoKilometers returns Kilometers in x Miles
func MilestoKilometers(kilometers float64) float64 {
	return kilometers * Kmmi

}

// KilometerstoMiles return Miles in x Kilometers
func KilometerstoMiles(miles float64) float64 {
	return miles / Kmmi
}

// FeetotoMeters converts feet to meters
func FeetotoMeters(feet float64) float64 {
	feet *= 0.3048
	return feet
}

// FeetoYard (feet float64) returns flaot32 in yard value
func FeetoYard(feet float64) float64 {
	return feet * 0.33333
}

// MeterstoFeet converts Meters to feet
func MeterstoFeet(meter float64) float64 {
	return meter * 3.2808
}

// MeterstoKilometer converts Meters to feet
func MeterstoKilometer(meter float64) float64 {
	return meter / 100
}

// MeterstoMile converts Meters to feet
func MeterstoMile(meter float64) float64 {
	return meter / 1609
}

// InchestoCentimeter (inches float64) returns float64 in Centimeter
func InchestoCentimeter(inches float64) float64 {
	return inches / 0.39370
}

// InchestoMillimeters (inches float64) returns float64 in Millimeters value
func InchestoMillimeters(inches float64) float64 {
	return InchestoCentimeter(inches) * 10
}

// InchestoFeet (inches float64) returns float64 in Feet value
func InchestoFeet(inches float64) float64 {
	return inches * 0.083333
}

// MeterstoYard (meters float64) returns float64 in Yard value
func MeterstoYard(meters float64) float64 {
	return meters * Yame
}

// YardtoMeters (yard float64) returns float64 in Meters value
func YardtoMeters(yard float64) float64 {
	return yard / Yame
}

// YardtoFeet (feet float64) returns float64 in Yard value
func YardtoFeet(yard float64) float64 {
	return yard / 3.0000
}

// AcretoHectare (acre float64) float64 in Hectare value
func AcretoHectare(acre float64) float64 {
	return acre / 2.471
}

// AcretoSquareInch (acre float64) float64 in Square inch value
func AcretoSquareInch(acre float64) float64 {
	return acre * 6.273e+6
}

// AcretoSquareFoot (acre float64) float64 in Square Foot value
func AcretoSquareFoot(acre float64) float64 {
	return acre * 43560
}

// AcretoSquareYard (acre float64) float64 in Square yard value
func AcretoSquareYard(acre float64) float64 {
	return acre * 4840
}

// AcretoSquareMile (acre float64) float64 in Square mile
func AcretoSquareMile(acre float64) float64 {
	return acre / 640
}

// AcretoSquareMeter (acre float64) float64 in Square meter
func AcretoSquareMeter(acre float64) float64 {
	return acre / 4046.856
}

// AcretoSquareKilometer (acre float64) float64 in Square mile
func AcretoSquareKilometer(acre float64) float64 {
	return acre / 247.105
}

// HectaretoAcre (hectare float64) float64 in Acre value
func HectaretoAcre(hectare float64) float64 {
	return hectare * 2.471
}

// HectaretoSquareInch (hectare float64) float64 in Square Inch value
func HectaretoSquareInch(hectare float64) float64 {
	return hectare * 1.55e+7
}

// HectaretoSquareFoot (hectare float64) float64 in Square Foot value
func HectaretoSquareFoot(hectare float64) float64 {
	return hectare * 107639.104
}

// HectaretoSquareYard (hectare float64) float64 in Square yard value
func HectaretoSquareYard(hectare float64) float64 {
	return hectare * 107639.104
}

// HectaretoSquareMile (hectare float64) float64 in Square mile value
func HectaretoSquareMile(hectare float64) float64 {
	return hectare / 258.999
}

// HectaretoSquareMeter (hectare float64) float64 in Square meter value
func HectaretoSquareMeter(hectare float64) float64 {
	return hectare * 10000
}

// HectaretoSquareKilometer (hectare float64) float64 in Square kilometer value
func HectaretoSquareKilometer(hectare float64) float64 {
	return hectare / 100
}

// SquareInchtoAcre (squareinch float64) float64 in Acre value
func SquareInchtoAcre(squareinch float64) float64 {
	return squareinch / 6.273e+6
}

// SquareInchtoHectare (squareinch float64) float64 in hectare value
func SquareInchtoHectare(squareinch float64) float64 {
	return squareinch / 1.55e+7
}

// SquareInchtoSquareFoot (squarefoot float64) float64 in Square Foot value
func SquareInchtoSquareFoot(squareinch float64) float64 {
	return squareinch / 144
}

// SquareInchtoSquareYard (squareinch float64) float64 in square yard value
func SquareInchtoSquareYard(squareinch float64) float64 {
	return squareinch / 1296
}

// SquareInchtoSquareMile (squareinch float64) float64 Square mile
func SquareInchtoSquareMile(squareinch float64) float64 {
	return squareinch / 4.014e+9
}

// SquareInchtoSquareMeter (squareinch float64) float64 Square meter value
func SquareInchtoSquareMeter(squareinch float64) float64 {
	return squareinch / 1550.003
}

// SquareInchtoSquareKilometer (squareinch float64) float64 Square kilometer value
func SquareInchtoSquareKilometer(squareinch float64) float64 {
	return squareinch / 1.55e+9
}

// SquareFoottoAcre (squareinch float64) float64 Square Acre value
func SquareFoottoAcre(squarefoot float64) float64 {
	return squarefoot / 43560
}

// SquareFoottoHectare (squareinch float64) float64 Square hectare value
func SquareFoottoHectare(squarefoot float64) float64 {
	return squarefoot / 107639.104
}

// SquareFoottoSquareInch (squarefoot float64) float64 Square Square inch value
func SquareFoottoSquareInch(squarefoot float64) float64 {
	return squarefoot * 144
}

// SquareFoottoSquareYard (squarefoot float64) float64 Square Square yard value
func SquareFoottoSquareYard(squarefoot float64) float64 {
	return squarefoot / 1296
}

// SquareFoottoSquareMile (squarefoot float64) float64 Square Square mile value
func SquareFoottoSquareMile(squarefoot float64) float64 {
	return squarefoot / 4.014e+9
}

// SquareFoottoSquareMeter (squareinch float64) float64 Square Square meter value
func SquareFoottoSquareMeter(squarefoot float64) float64 {
	return squarefoot / 1550.003
}

// SquareFoottoSquareKilometer (squareinch float64) float64 Square Square kilometer value
func SquareFoottoSquareKilometer(squarefoot float64) float64 {
	return squarefoot / 1.55e+9
}

// SquareYardtoAcre (squareyard float64) float64 acre in value
func SquareYardtoAcre(squareyard float64) float64 {
	return squareyard / 4840
}

// SquareYardtoHectare (squareyard float64) float64 acre in Hectare value
func SquareYardtoHectare(squareyard float64) float64 {
	return squareyard / 11959.9
}

// SquareYardtoSquareInch (squareyard float64) float64 acre in Square inch value
func SquareYardtoSquareInch(squareyard float64) float64 {
	return squareyard * 1296
}

// SquareYardtoSquareFoot (squareyard float64) float64 acre in Square foot value
func SquareYardtoSquareFoot(squareyard float64) float64 {
	return squareyard * 9
}

// SquareYardtoSquareMile (squareyard float64) float64 acre in Square mile value
func SquareYardtoSquareMile(squareyard float64) float64 {
	return squareyard / 3.098e+6
}

// SquareYardtoSquareMeter (squareyard float64) float64 acre in Square meter value
func SquareYardtoSquareMeter(squareyard float64) float64 {
	return squareyard / 1.196
}

// SquareYardtoSquareKilometer (squareyard float64) float64 acre in Square kilometer value
func SquareYardtoSquareKilometer(squareyard float64) float64 {
	return squareyard / 1.196e+6
}

//SquareMiletoAcre (squaremile float64) float64
func SquareMiletoAcre(squaremile float64) float64 {
	return squaremile * 640
}

//SquareMiletoHectare (squaremile float64) float64
func SquareMiletoHectare(squaremile float64) float64 {
	return squaremile * 258.999
}

//SquareMiletoSquareInch (squaremile float64) float64
func SquareMiletoSquareInch(squaremile float64) float64 {
	return squaremile * 4.014e+9
}

//SquareMiletoSquareFoot (squaremile float64) float64
func SquareMiletoSquareFoot(squaremile float64) float64 {
	return squaremile * 2.788e+7
}

// SquareMiletoSquareYard (squaremile float64) float64
func SquareMiletoSquareYard(squaremile float64) float64 {
	return squaremile * 3.098e+6
}

//SquareMiletometer (squaremile float64) float64
func SquareMiletometer(squaremile float64) float64 {
	return squaremile * 2.59e+6
}

//SquareMiletokilometer (squaremile float64) float64
func SquareMiletokilometer(squaremile float64) float64 {
	return squaremile * 2.59
}

// SquareKilometertoAcre (squarekilometer float64) float64
func SquareKilometertoAcre(squarekilometer float64) float64 {
	return squarekilometer * 247.105
}

// SquareKilometertoHectare (squarekilometer float64) float64
func SquareKilometertoHectare(squarekilometer float64) float64 {
	return squarekilometer * 100
}

// SquareKilometertoSquareInch (squarekilometer float64) float64
func SquareKilometertoSquareInch(squarekilometer float64) float64 {
	return squarekilometer * 1.55e+9
}

// SquareKilometertoSquareFoot (squarekilometer float64) float64
func SquareKilometertoSquareFoot(squarekilometer float64) float64 {
	return squarekilometer * 1.076e+7
}

// SquareKilometertoSquareYard (squarekilometer float64) float64
func SquareKilometertoSquareYard(squarekilometer float64) float64 {
	return squarekilometer * 1.196e+6
}

// SquareKilometertoSquareMile (squarekilometer float64) float64
func SquareKilometertoSquareMile(squarekilometer float64) float64 {
	return squarekilometer / 2.59
}

// SquareKilometertoSquareMeter (squarekilometer float64) float64
func SquareKilometertoSquareMeter(squarekilometer float64) float64 {
	return squarekilometer * 1e+6
}
