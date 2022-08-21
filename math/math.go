package math

import "math"

func MyRoundFloat(val float64, precision int8) float64 {
	ratio := math.Pow10(int(precision))
	return math.Round(val*ratio) / ratio
}

func MyRoundFloat2(f float64, digits int) float64 {
	if math.Abs(f) < 0.5 {
		return 0
	}
	pow := math.Pow10(digits)
	return math.Trunc(f*pow+math.Copysign(0.5, f)) / pow
}
