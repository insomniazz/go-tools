package math

import "math"

func MyRoundFloat(val float64, precision int8) float64 {
	ratio := math.Pow10(int(precision))
	return math.Round(val*ratio) / ratio
}
