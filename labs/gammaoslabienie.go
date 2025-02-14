package labs

import (
	PiWo "PiWo/src"
	fmt "fmt"
	math "math"
)

func GammaOslabienie() {
	PiWo.Start(6, "Osłabienie promieniowania gamma")

	X_1 := PiWo.SI([]float64{2, 5, 7.1, 10.05, 12.2, 15.1, 16.9, 20}, 1e-3)
	n_1 := []float64{886, 737, 657, 615, 497, 440, 441, 398}
	uN_1 := []float64{}
	N_1 := []float64{}

	X_2 := PiWo.SI([]float64{2, 5, 6.85, 9.9, 11.85, 14.75, 16.85, 19.9}, 1e-3)
	n_2 := []float64{772, 608, 535, 457, 398, 369, 315, 230}
	uN_2 := []float64{}
	N_2 := []float64{}

	X_3 := PiWo.SI([]float64{4.8, 9.75, 14.65, 19.85}, 1e-3)
	n_3 := []float64{772, 684, 674, 653}
	uN_3 := []float64{}
	N_3 := []float64{}

	for i := range n_1 {
		N_1 = append(N_1, math.Log(n_1[i]))
	}
	for i := range N_1 {
		uN_1 = append(uN_1, math.Sqrt(n_1[i])/n_1[i])
	}

	for i := range n_2 {
		N_2 = append(N_2, math.Log(n_2[i]))
	}
	for i := range N_2 {
		uN_2 = append(uN_2, math.Sqrt(n_2[i])/n_2[i])
	}

	for i := range n_3 {
		N_3 = append(N_3, math.Log(n_3[i]))
	}
	for i := range N_3 {
		uN_3 = append(uN_3, math.Sqrt(n_3[i])/n_3[i])
	}

	SQ_1 := PiWo.SquaresAB(X_1, N_1, uN_1)
	SQ_2 := PiWo.SquaresAB(X_2, N_2, uN_2)
	SQ_3 := PiWo.SquaresAB(X_3, N_3, uN_3)

	fmt.Printf("u1 = %0.4f (%0.6f)\n", -SQ_1.A/1000, SQ_1.UA/1000)
	fmt.Printf("u2 = %0.4f (%0.6f)\n", -SQ_2.A/1000, SQ_2.UA/1000)
	fmt.Printf("u3 = %0.4f (%0.6f)\n", -SQ_3.A/1000, SQ_3.UA/1000)

	PiWo.Excel("GammaOslabienie", [][]float64{N_1, X_1})
}
