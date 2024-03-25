package labs

import (
	PiWo "PiWo/src"
	math "math"

	fmt "fmt"
)

func Dielektryk() {
	fmt.Println("\u001B[33mDoświadczenie 2:\u001B[0m Odbicie światła od powierzchni dielektryków")
	X := []float64{PiWo.Sin(3), PiWo.Sin(6), PiWo.Sin(10), PiWo.Sin(13), PiWo.Sin(16), PiWo.Sin(19), PiWo.Sin(22), PiWo.Sin(25), PiWo.Sin(28), PiWo.Sin(31), PiWo.Sin(33), PiWo.Sin(35), PiWo.Sin(37), PiWo.Sin(39), PiWo.Sin(40)}
	Y := []float64{PiWo.Sin(5), PiWo.Sin(10), PiWo.Sin(15), PiWo.Sin(20), PiWo.Sin(25), PiWo.Sin(30), PiWo.Sin(35), PiWo.Sin(40), PiWo.Sin(45), PiWo.Sin(50), PiWo.Sin(55), PiWo.Sin(60), PiWo.Sin(65), PiWo.Sin(70), PiWo.Sin(75)}
	U := []float64{math.Pi / 180, math.Pi / 180, math.Pi / 180, math.Pi / 180, math.Pi / 180, math.Pi / 180, math.Pi / 180, math.Pi / 180, math.Pi / 180, math.Pi / 180, math.Pi / 180, math.Pi / 180, math.Pi / 180, math.Pi / 180, math.Pi / 180}

	for i, y := range Y {
		U[i] *= PiWo.Cos(y)
	}

	SQ := PiWo.SquaresAB(X, Y, U)
	fmt.Printf("n = %0.2f (%0.2f) - pojedyńczy pomiar\n", PiWo.Sin(30)/PiWo.Sin(19), math.Sqrt(math.Pow(math.Pi/(180*PiWo.Cos(19)), 2)+math.Pow((math.Pi*PiWo.Sin(30)*PiWo.Cos(19))/(180*PiWo.Sin(19)*PiWo.Sin(19)), 2)))
	fmt.Printf("n = %0.2f (%0.2f) - metoda najmniejszych kwadratów\n", SQ.A, SQ.UA)
	fmt.Printf("n = %0.2f (%0.2f) - kąt graniczny\n", 1/PiWo.Sin(43), (math.Pi*math.Cos(43))/(180*math.Pow(PiWo.Sin(43), 2)))
	fmt.Printf("n = %0.2f (%0.2f) - kąt Brewstera\n", PiWo.Tg(56), (math.Pi)/(180*math.Pow(PiWo.Cos(56), 2)))

	fmt.Println()
}
