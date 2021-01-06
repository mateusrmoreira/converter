package converter

// With area interface you don't need to remember each function created to run for area, just let it run
// Function Area selects which unit the uniconv.go // file stands for

import (
	"formulas"
	"strings"
)

// AcreFuncs returns results of acre functions
func AcreFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "ya"):
		return formulas.AcretoSquareYard(value)
	case strings.Contains(unit, "ft"):
		return formulas.AcretoSquareFoot(value)
	case strings.Contains(unit, "in"):
		return formulas.AcretoSquareInch(value)
	case strings.Contains(unit, "hc"):
		return formulas.AcretoHectare(value)
	case strings.Contains(unit, "ml"):
		return formulas.AcretoSquareMile(value)
	case strings.Contains(unit, "km"):
		return formulas.AcretoSquareKilometer(value)
	case strings.Contains(unit, "m"):
		return formulas.AcretoSquareMeter(value)
	}
	return 0.0
}

// MetersFuncs (uni string, value float64) float64
func MetersFuncs(uni string, value float64) float64 {
	switch {
	case strings.Contains(uni, "ya"):
		return formulas.MeterstoYard(value)
	case strings.Contains(uni, "ft"):
		return formulas.MeterstoFeet(value)
	case strings.Contains(uni, "km"):
		return formulas.MeterstoKilometer(value)
	case strings.Contains(uni, "ml"):
		return formulas.MeterstoMile(value)

	}
	return 0.0
}

// FootFuncs returns the area
func FootFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "ya"):
		return formulas.SquareFoottoSquareYard(value)
	case strings.Contains(unit, "acre"):
		return formulas.SquareFoottoAcre(value)
	case strings.Contains(unit, "hc"):
		return formulas.SquareFoottoHectare(value)
	case strings.Contains(unit, "in"):
		return formulas.SquareFoottoSquareInch(value)
	case strings.Contains(unit, "ml"):
		return formulas.SquareFoottoSquareMile(value)
	case strings.Contains(unit, "m"):
		return formulas.SquareFoottoSquareMeter(value)
	case strings.Contains(unit, "ml"):
		return formulas.SquareFoottoSquareMile(value)
	case strings.Contains(unit, "km"):
		return formulas.SquareFoottoSquareKilometer(value)
	}
	return 0.0
}

// HectareFuncs (unit string, value float64) float64
func HectareFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "acre"):
		return formulas.HectaretoAcre(value)
	case strings.Contains(unit, "in"):
		return formulas.HectaretoSquareInch(value)
	case strings.Contains(unit, "ft"):
		return formulas.HectaretoSquareFoot(value)
	case strings.Contains(unit, "ya"):
		return formulas.HectaretoSquareYard(value)
	case strings.Contains(unit, "ml"):
		return formulas.HectaretoSquareMile(value)
	case strings.Contains(unit, "m"):
		return formulas.HectaretoSquareMeter(value)
	case strings.Contains(unit, "km"):
		return formulas.HectaretoSquareKilometer(value)
	}
	return 0.0
}

// InchFuncs (unit string, value float64) float64
func InchFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "acre"):
		return formulas.SquareInchtoAcre(value)
	case strings.Contains(unit, "hc"):
		return formulas.SquareInchtoHectare(value)
	case strings.Contains(unit, "ft"):
		return formulas.SquareInchtoSquareFoot(value)
	case strings.Contains(unit, "ya"):
		return formulas.SquareInchtoSquareYard(value)
	case strings.Contains(unit, "ml"):
		return formulas.SquareInchtoSquareMile(value)
	case strings.Contains(unit, "m"):
		return formulas.SquareInchtoSquareMeter(value)
	case strings.Contains(unit, "km"):
		return formulas.SquareInchtoSquareKilometer(value)
	}
	return 0.0
}

// KilometerFuncs returns the results based on KM
func KilometerFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "acre"):
		return formulas.SquareKilometertoAcre(value)
	case strings.Contains(unit, "hc"):
		return formulas.SquareKilometertoHectare(value)
	case strings.Contains(unit, "ft"):
		return formulas.SquareKilometertoSquareFoot(value)
	case strings.Contains(unit, "ya"):
		return formulas.SquareKilometertoSquareYard(value)
	case strings.Contains(unit, "ml"):
		return formulas.SquareKilometertoSquareMile(value)
	case strings.Contains(unit, "m"):
		return formulas.SquareKilometertoSquareMeter(value)
	case strings.Contains(unit, "in"):
		return formulas.SquareKilometertoSquareInch(value)
	}
	return 0.0
}

// YardFuncs (unit string, value float64) float64
func YardFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "acre"):
		return formulas.SquareYardtoAcre(value)
	case strings.Contains(unit, "hc"):
		return formulas.SquareYardtoHectare(value)
	case strings.Contains(unit, "ft"):
		return formulas.SquareYardtoSquareFoot(value)
	case strings.Contains(unit, "km"):
		return formulas.SquareYardtoSquareKilometer(value)
	case strings.Contains(unit, "ml"):
		return formulas.SquareYardtoSquareMile(value)
	case strings.Contains(unit, "m"):
		return formulas.SquareYardtoSquareMeter(value)
	case strings.Contains(unit, "in"):
		return formulas.SquareYardtoSquareInch(value)
	}
	return 0.0
}

// MileFuncs (unit string, value float64) float64
func MileFuncs(unit string, value float64) float64 {
	switch {
	case strings.Contains(unit, "acre"):
		return formulas.SquareYardtoAcre(value)
	case strings.Contains(unit, "hc"):
		return formulas.SquareYardtoHectare(value)
	case strings.Contains(unit, "ft"):
		return formulas.SquareYardtoSquareFoot(value)
	case strings.Contains(unit, "km"):
		return formulas.SquareYardtoSquareKilometer(value)
	case strings.Contains(unit, "ml"):
		return formulas.SquareYardtoSquareMile(value)
	case strings.Contains(unit, "m"):
		return formulas.SquareYardtoSquareMeter(value)
	case strings.Contains(unit, "in"):
		return formulas.SquareYardtoSquareInch(value)
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
