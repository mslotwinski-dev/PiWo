package labs

import (
	PiWo "PiWo/src"
	fmt "fmt"
	math "math"
)

func Interferencja() {
	PiWo.Start(5, "Pomiar długości fali elektromagnetycznej")
	ua := (math.Pi / 180) / math.Sqrt(3)
	ux := 1e-3 / math.Sqrt(3)

	fmt.Println("Siatka dyfracyjna")

	l := 88.1e-2
	n := 12.0

	d := l / n
	// ud := ux / n

	l1 := d * PiWo.Sin(53) / 2
	l2 := d * PiWo.Sin(24)
	l3 := d * PiWo.Sin(26)
	l4 := d * PiWo.Sin(51) / 2

	fmt.Printf("λ1 = %fm\n", l1)
	fmt.Printf("λ2 = %fm\n", l2)
	fmt.Printf("λ3 = %fm\n", l3)
	fmt.Printf("λ4 = %fm\n", l4)

	ual := PiWo.SDevMean([]float64{l1, l2, l3, l4})

	fmt.Printf("λ = %.2f (%.2f) mm\n", 1000*PiWo.Mean([]float64{l1, l2, l3, l4}), 1000*ual)

	fmt.Println("Michaelson")
	x0 := 29.0
	xn := 302.0
	fmt.Printf("λ = %.2f (%.2f) mm\n", 2*(xn-x0)/16, 2000*PiWo.SumSquares([]float64{ux, ux})/16)

	fmt.Println("FP z kątem")
	deltaxi0 := []float64{101, 88, 68, 54, 34, 18}
	i := []float64{1, 2, 3, 4, 5, 6}
	udeltax := []float64{ux, ux, ux, ux, ux, ux}

	SQ := PiWo.SquaresAB(i, deltaxi0, udeltax)

	fmt.Printf("λ = %.2f (%.2f) mm\n", 2*PiWo.Cos(20)*SQ.A, 2*PiWo.SumSquares([]float64{SQ.UA * PiWo.Cos(20), SQ.A * PiWo.Sin(20) * ua}))

	fmt.Println("FP bez kąta")
	fmt.Printf("λ = %.2f (%.2f) mm\n", 2.0*244/15, 2000*PiWo.SumSquares([]float64{ux, ux})/15)

}
