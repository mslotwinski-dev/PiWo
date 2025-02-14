package main

import (
	// PiWo "PiWo/src"
	Lab "PiWo/labs"
	fmt "fmt"
)

func main() {
	fmt.Println()
	fmt.Println("\u001B[36mPiWo\u001B[0m - \u001B[33mty pijesz piwo, a ja licze ci sprawozdania\u001B[0m")
	fmt.Println("by \u001B[31mMateusz Słotwiński\u001B[0m")
	fmt.Println()

	Lab.DeBroile()
	Lab.Dielektryk()
	Lab.Laminarny()
	Lab.DrganiaRLC()
	Lab.Interferencja()
	Lab.GammaOslabienie()
	Lab.Magnetyczne()
}
