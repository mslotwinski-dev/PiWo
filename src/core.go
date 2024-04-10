package src

import (
	fmt "fmt"
)

func Start(nr int, title string) {
	fmt.Printf("\u001B[33mDoświadczenie %d:\u001B[0m %s\n", nr, title)
}
