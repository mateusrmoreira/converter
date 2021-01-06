package converter

// With area interface you don't need to remember each function created to run for area, just let it run
// Function Area selects which unit the uniconv.go // file stands for

import (
	"strings"
)

// AcreFuncs returns results of acre functions
func AcreFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "ya"):
		return AcretoSquareYard(value)
	case strings.Contains(unit, "ft"):
		return AcretoSquareFoot(value)
	case strings.Contains(unit, "in"):
		return AcretoSquareInch(value)
	case strings.Contains(unit, "hc"):
		return AcretoHectare(value)
	case strings.Contains(unit, "ml"):
		return AcretoSquareMile(value)
	case strings.Contains(unit, "km"):
		return AcretoSquareKilometer(value)
	case strings.Contains(unit, "m"):
		return AcretoSquareMeter(value)
	}
	return 0.0
}

// MetersFuncs (uni string, value float64) float64
func MetersFuncs(uni string, value float64) float64 {
	switch {
	case strings.Contains(uni, "ya"):
		return MeterstoYard(value)
	case strings.Contains(uni, "ft"):
		return MeterstoFeet(value)
	case strings.Contains(uni, "km"):
		return MeterstoKilometer(value)
	case strings.Contains(uni, "ml"):
		return MeterstoMile(value)

	}
	return 0.0
}

// FootFuncs returns the area
func FootFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "ya"):
		return SquareFoottoSquareYard(value)
	case strings.Contains(unit, "acre"):
		return SquareFoottoAcre(value)
	case strings.Contains(unit, "hc"):
		return SquareFoottoHectare(value)
	case strings.Contains(unit, "in"):
		return SquareFoottoSquareInch(value)
	case strings.Contains(unit, "ml"):
		return SquareFoottoSquareMile(value)
	case strings.Contains(unit, "m"):
		return SquareFoottoSquareMeter(value)
	case strings.Contains(unit, "ml"):
		return SquareFoottoSquareMile(value)
	case strings.Contains(unit, "km"):
		return SquareFoottoSquareKilometer(value)
	}
	return 0.0
}

// HectareFuncs (unit string, value float64) float64
func HectareFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "acre"):
		return HectaretoAcre(value)
	case strings.Contains(unit, "in"):
		return HectaretoSquareInch(value)
	case strings.Contains(unit, "ft"):
		return HectaretoSquareFoot(value)
	case strings.Contains(unit, "ya"):
		return HectaretoSquareYard(value)
	case strings.Contains(unit, "ml"):
		return HectaretoSquareMile(value)
	case strings.Contains(unit, "m"):
		return HectaretoSquareMeter(value)
	case strings.Contains(unit, "km"):
		return HectaretoSquareKilometer(value)
	}
	return 0.0
}

// InchFuncs (unit string, value float64) float64
func InchFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "acre"):
		return SquareInchtoAcre(value)
	case strings.Contains(unit, "hc"):
		return SquareInchtoHectare(value)
	case strings.Contains(unit, "ft"):
		return SquareInchtoSquareFoot(value)
	case strings.Contains(unit, "ya"):
		return SquareInchtoSquareYard(value)
	case strings.Contains(unit, "ml"):
		return SquareInchtoSquareMile(value)
	case strings.Contains(unit, "m"):
		return SquareInchtoSquareMeter(value)
	case strings.Contains(unit, "km"):
		return SquareInchtoSquareKilometer(value)
	}
	return 0.0
}

// KilometerFuncs returns the results based on KM
func KilometerFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "acre"):
		return SquareKilometertoAcre(value)
	case strings.Contains(unit, "hc"):
		return SquareKilometertoHectare(value)
	case strings.Contains(unit, "ft"):
		return SquareKilometertoSquareFoot(value)
	case strings.Contains(unit, "ya"):
		return SquareKilometertoSquareYard(value)
	case strings.Contains(unit, "ml"):
		return SquareKilometertoSquareMile(value)
	case strings.Contains(unit, "m"):
		return SquareKilometertoSquareMeter(value)
	case strings.Contains(unit, "in"):
		return SquareKilometertoSquareInch(value)
	}
	return 0.0
}

// YardFuncs (unit string, value float64) float64
func YardFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "acre"):
		return SquareYardtoAcre(value)
	case strings.Contains(unit, "hc"):
		return SquareYardtoHectare(value)
	case strings.Contains(unit, "ft"):
		return SquareYardtoSquareFoot(value)
	case strings.Contains(unit, "km"):
		return SquareYardtoSquareKilometer(value)
	case strings.Contains(unit, "ml"):
		return SquareYardtoSquareMile(value)
	case strings.Contains(unit, "m"):
		return SquareYardtoSquareMeter(value)
	case strings.Contains(unit, "in"):
		return SquareYardtoSquareInch(value)
	}
	return 0.0
}

// MileFuncs (unit string, value float64) float64
func MileFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "acre"):
		return SquareYardtoAcre(value)
	case strings.Contains(unit, "hc"):
		return SquareYardtoHectare(value)
	case strings.Contains(unit, "ft"):
		return SquareYardtoSquareFoot(value)
	case strings.Contains(unit, "km"):
		return SquareYardtoSquareKilometer(value)
	case strings.Contains(unit, "ml"):
		return SquareYardtoSquareMile(value)
	case strings.Contains(unit, "m"):
		return SquareYardtoSquareMeter(value)
	case strings.Contains(unit, "in"):
		return SquareYardtoSquareInch(value)
	}
	return 0.0
}

//Area functions that calls an area
func Area(unit1, unit2 string, value float64) float64 {
	switch {
	case strings.Contains(unit1, "m"):
		return MetersFuncs(unit2, value)
	case strings.Contains(unit1, "acre"):
		return AcreFuncs(unit2, value)
	case strings.Contains(unit1, "ft"):
		return FootFuncs(unit2, value)
	case strings.Contains(unit1, "hc"):
		return HectareFuncs(unit2, value)
	case strings.Contains(unit1, "in"):
		return InchFuncs(unit2, value)
	case strings.Contains(unit1, "km"):
		return KilometerFuncs(unit2, value)
	case strings.Contains(unit1, "ya"):
		return YardFuncs(unit2, value)
	case strings.Contains(unit1, "ml"):
		return MileFuncs(unit2, value)
	}
	return 0.0
}
