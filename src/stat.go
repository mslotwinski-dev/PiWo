package src

import (
	fmt "fmt"
	"math"
)

func Init() {
	fmt.Println("Huj")
}

func SI(arr []float64, i float64) []float64 {
	for n := range arr {
		arr[n] *= i
	}
	return arr
}

type Squares struct {
	A  float64
	UA float64
	B  float64
	UB float64
}

func SquaresA(x []float64, y []float64, uy []float64) Squares {
	l := 0.0
	m := 0.0

	for i := range x {
		l += (x[i] * y[i]) / math.Pow(uy[i], 2)
		m += math.Pow(x[i], 2) / math.Pow(uy[i], 2)
	}

	a := l / m
	ua := math.Sqrt(1 / m)

	return Squares{a, ua, 0, 0}
}

func SquaresAB(x []float64, y []float64, uy []float64) Squares {

	S_1 := 0.0 // W
	S_2 := 0.0 // WXY
	S_3 := 0.0 // WX
	S_4 := 0.0 // WY
	S_5 := 0.0 // WX2

	for i := range x {
		w_i := 1 / math.Pow(uy[i], 2)

		S_1 += w_i
		S_2 += w_i * x[i] * y[i]
		S_3 += w_i * x[i]
		S_4 += w_i * y[i]
		S_5 += w_i * x[i] * x[i]

	}

	a := (S_1*S_2 - S_3*S_4) / (S_1*S_5 - math.Pow(S_3, 2))
	ua := math.Sqrt((S_1) / (S_1*S_5 - math.Pow(S_3, 2)))
	b := (S_5*S_4 - S_3*S_2) / (S_1*S_5 - math.Pow(S_2, 2))
	ub := math.Sqrt(((S_5) / (S_1*S_5 - math.Pow(S_2, 2))))

	return Squares{a, ua, b, ub}
}
