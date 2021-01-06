package converter

import (
	"strings"
)

// FootSpeedFuncs returns the foot speed in others metric systems
func FootSpeedFuncs(uni string, value float64) float64 {
	switch {
	case strings.Contains(uni, "speed_no"):
		return FootoKnot(value)
	case strings.Contains(uni, "speed_km"):
		return FootoKilometerh(value)
	case strings.Contains(uni, "speed_mi"):
		return FootoMilesh(value)
	}
	return 0.0
}

// MeterSpeedFuncs returns the Meter speed in others metric systems
func MeterSpeedFuncs(uni string, value float64) float64 {
	switch {
	case strings.Contains(uni, "speed_no"):
		return MeterstoKnot(value)
	case strings.Contains(uni, "speed_km"):
		return MeterstoKilometerh(value)
	case strings.Contains(uni, "speed_mi"):
		return MeterstoMilesh(value)
	case strings.Contains(uni, "speed_ft"):
		return MeterstoFoot(value)
	}
	return 0.0
}

// Speed return the output from all speed functions
func Speed(unit1, unit2 string, value float64) float64 {
	switch {
	case strings.Contains(unit1, "speed_ft"):
		return FootSpeedFuncs(unit2, value)
	case strings.Contains(unit1, "speed_m"):
		return MeterSpeedFuncs(unit2, value)
	case strings.Contains(unit1, "speed_km"):
		return 0.0
	case strings.Contains(unit1, "speed_mi"):
		return 0.0
	case strings.Contains(unit1, "speed_no"):
		return 0.0
	}
	return 0.0
}
