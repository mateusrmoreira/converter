package formulas

// KnottoKilomerh (knot float64) float64
func KnottoKilomerh(knot float64) float64 {
	return knot * 1.852
}

// KnottoMeterS (knot float64) float64
func KnottoMeterS(knot float64) float64 {
	return knot / 1.944
}

// KnottoFeetS (knot float64) float64
func KnottoFeetS(knot float64) float64 {
	return knot * 1.688
}

// KnottoMilesh (knot float64) float64
func KnottoMilesh(knot float64) float64 {
	return knot * 1.151
}

// MileshtoKilometerh (milesh float64) float64
func MileshtoKilometerh(milesh float64) float64 {
	return milesh * 1.609
}

// MileshtoMeters (milesh float64) float64
func MileshtoMeters(milesh float64) float64 {
	return milesh / 2.237
}

// MileshtoFoot (milesh float64) float64
func MileshtoFoot(milesh float64) float64 {
	return milesh * 1.467
}

// MileshtoKnot (milesh float64) float64
func MileshtoKnot(milesh float64) float64 {
	return milesh * 1.51
}

// KilometerhtoKnot (Kilometerh float64) float64
func KilometerhtoKnot(Kilometerh float64) float64 {
	return Kilometerh / 1.852
}

// KilometerhtoMeters (Kilometerh float64) float64
func KilometerhtoMeters(Kilometerh float64) float64 {
	return Kilometerh / 3.6
}

// KilometerhtoMilesh (Kilometerh float64) float64
func KilometerhtoMilesh(Kilometerh float64) float64 {
	return Kilometerh / 1.609
}

// KilometerhtoFeet (Kilometerh float64) float64
func KilometerhtoFeet(Kilometerh float64) float64 {
	return Kilometerh / 1.097
}

// MeterstoMilesh (meters float64) float64
func MeterstoMilesh(meters float64) float64 {
	return meters * 2.237
}

// MeterstoKnot (meters float64) float64
func MeterstoKnot(meters float64) float64 {
	return meters * 1.944
}

// MeterstoKilometerh (meters float64) float64
func MeterstoKilometerh(meters float64) float64 {
	return meters * 3.6
}

// MeterstoFoot (meters float64) float64
func MeterstoFoot(meters float64) float64 {
	return meters * 3.281
}

// FootoKnot (foots float64) float64
func FootoKnot(foots float64) float64 {
	return foots / 1.688
}

// FootoKilometerh (foots float64) float64
func FootoKilometerh(foots float64) float64 {
	return foots * 1.097
}

// FootoMeters (foots float64) float64
func FootoMeters(foots float64) float64 {
	return foots / 3.281
}

// FootoMilesh (foots float64) float64
func FootoMilesh(foots float64) float64 {
	return foots / 1.467
}
