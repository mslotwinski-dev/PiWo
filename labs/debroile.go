package labs

import (
	PiWo "PiWo/src"

	fmt "fmt"
	math "math"
)

func DeBroile() {
	fmt.Println("\u001B[33mDoświadczenie 1:\u001B[0m Fale De Broile'a")

	const r = 127e-3

	U_1 := PiWo.SI([]float64{200, 230, 260, 289, 321, 352, 381, 407, 439, 530}, 7.5) // V * 7,5
	D_11 := PiWo.SI([]float64{43.2, 40.6, 38.0, 35.7, 34.4, 32.6, 30.6, 29.7, 29.0, 26.0}, 1e-3)
	D_12 := PiWo.SI([]float64{73.2, 69.6, 65.1, 60.9, 57.3, 55.8, 52.1, 50.7, 48.7, 44.6}, 1e-3)
	U_D1 := PiWo.SI([]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05}, 1e-3/math.Sqrt(3))

	U_2 := PiWo.SI([]float64{2.3, 2.6, 2.9, 3.15, 3.30, 3.45, 3.60, 3.75, 4.0, 4.3, 4.6, 4.9, 5.34, 5.6, 6}, 1000) // V * 7,5
	D_21 := PiWo.SI([]float64{16, 15, 14, 14, 13, 12, 12, 12, 11, 11, 11, 10, 10, 10, 9}, 2e-3)
	D_22 := PiWo.SI([]float64{28, 27, 24, 23, 23, 23, 22, 21, 21, 20, 19, 18, 17, 17, 17}, 2e-3)
	U_D2 := PiWo.SI([]float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1e-3/math.Sqrt(3))

	X_1 := make([]float64, len(U_1))
	copy(X_1, U_1)
	X_2 := make([]float64, len(U_2))
	copy(X_2, U_2)

	for i := range X_1 {
		X_1[i] = 1 / math.Sqrt(X_1[i])
	}

	for i := range X_2 {
		X_2[i] = 1 / math.Sqrt(X_2[i])
	}

	SQ_11 := PiWo.SquaresAB(X_1, D_11, U_D1)
	SQ_12 := PiWo.SquaresAB(X_1, D_12, U_D1)
	SQ_21 := PiWo.SquaresAB(X_2, D_21, U_D2)
	SQ_22 := PiWo.SquaresAB(X_2, D_22, U_D2)

	d_1 := (r * h / SQ_11.A) * math.Sqrt(2/(m_e*e))
	d_2 := (2 * r * h / SQ_12.A) * math.Sqrt(2/(m_e*e))
	d_3 := (r * h / SQ_21.A) * math.Sqrt(2/(m_e*e))
	d_4 := (2 * r * h / SQ_22.A) * math.Sqrt(2/(m_e*e))

	ud_1 := (r*h*SQ_11.UA/(math.Pow(SQ_11.A, 2)))*math.Sqrt(2/(m_e*e)) + (h*0.0003/(math.Pow(SQ_11.A, 1)))*math.Sqrt(2/(m_e*e))
	ud_2 := (2*r*h*SQ_12.UA/(math.Pow(SQ_12.A, 2)))*math.Sqrt(2/(m_e*e)) + (2*h*0.0003/(math.Pow(SQ_12.A, 1)))*math.Sqrt(2/(m_e*e))
	ud_3 := (r*h*SQ_21.UA/(math.Pow(SQ_21.A, 2)))*math.Sqrt(2/(m_e*e)) + (h*0.0003/(math.Pow(SQ_21.A, 1)))*math.Sqrt(2/(m_e*e))
	ud_4 := (2*r*h*SQ_22.UA/(math.Pow(SQ_22.A, 2)))*math.Sqrt(2/(m_e*e)) + (2*h*0.0003/(math.Pow(SQ_22.A, 1)))*math.Sqrt(2/(m_e*e))

	fmt.Printf("d1 = %0.2f (%0.2f)\n", d_1*1e12, ud_1*1e12)
	fmt.Printf("d2 = %0.2f (%0.2f)\n", d_2*1e12, ud_2*1e12)
	fmt.Printf("d3 = %0.2f (%0.2f)\n", d_3*1e12, ud_3*1e12)
	fmt.Printf("d4 = %0.2f (%0.2f)\n", d_4*1e12, ud_4*1e12)

	// PiWo.PlotSquares("DeBroile", "U^-1/2", "D_11", U_1, D_11, SQ_11)
	// PiWo.PlotSquares("DeBroile", "U^-1/2", "D_12", U_1, D_12, SQ_12)
	// PiWo.PlotSquares("DeBroile", "U^-1/2", "D_21", U_2, D_21, SQ_21)
	// PiWo.PlotSquares("DeBroile", "U^-1/2", "D_22", U_2, D_22, SQ_22)

	PiWo.Excel("DeBroile", [][]float64{U_1, D_11, D_12, U_D1, X_1, U_2, D_21, D_22, U_D2, X_2})

	// fmt.Println(h * 10e12 / math.Sqrt(2*e*m_e*6000))

	fmt.Println()
}
