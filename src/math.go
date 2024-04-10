package src

import (
	"math"
)

func SI(arr []float64, i float64) []float64 {
	for n := range arr {
		arr[n] *= i
	}
	return arr
}

func Sin(x float64) float64 {
	return math.Sin(x * math.Pi / 180)
}

func Cos(x float64) float64 {
	return math.Cos(x * math.Pi / 180)
}

func Tg(x float64) float64 {
	return Sin(x) / Cos(x)
}

func SumSquares(x []float64) float64 {
	sum := 0.0
	for _, v := range x {
		sum += math.Pow(v, 2)
	}

	return math.Sqrt(sum)

}
