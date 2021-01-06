package converter

// variables to parse in the calc
var (
	area = []string{
		"area",
		"acre",
		"ft",
		"hc",
		"in",
		"km",
		"m",
		"ml",
		"km",
		"ya",
	}

	time = []string{
		"time",
		"century",
		"decade",
		"day",
		"hour",
		"sec",
		"min",
		"month",
		"week",
		"year",
	}
	temperature = []string{
		"temperature",
		"c",
		"f",
		"k",
	}

	speed = []string{
		"speed",
		"speed_ft",
		"speed_km",
		"speed_m",
		"speed_mi",
		"speed_no",
	}
)

func contains(unitsindex []string, measure string) bool {
	for _, value := range unitsindex {
		if measure == value {
			return true
		}
	}
	return false
}

// Parser makes a par to define which unit the user refers
// metrics string is used to identify the kind of metric it will be used
func Parser(unit1, unit2 string, value float64) float64 {
	switch {
	case contains(area, unit1) && contains(area, unit2):
		result := Area(unit1, unit2, value)
		return result
	case contains(speed, unit1) && contains(speed, unit2):
		result := Speed(unit1, unit2, value)
		return result
	case contains(temperature, unit1) && contains(temperature, unit2):
		result := Temperature(unit1, unit2, value)
		return result
	case contains(time, unit1) && contains(time, unit2):
		result := Time(unit1, unit2, value)
		return result
	default:
		return 0.0
	}
	return 0.0
}
