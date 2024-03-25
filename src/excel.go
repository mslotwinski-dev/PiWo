package src

import (
	fmt "fmt"

	"github.com/xuri/excelize/v2"
)

func Excel(name string, Data [][]float64) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	index, err := f.NewSheet("PiWo")
	if err != nil {
		fmt.Println(err)
		return
	}

	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for i, column := range Data {

		f.SetCellValue("PiWo", letters[i]+"1", letters[i])
		for j, v := range column {
			f.SetCellValue("PiWo", letters[i]+fmt.Sprint(j+2), v)
		}

	}

	f.SetActiveSheet(index)
	if err := f.SaveAs("assets/excel/" + name + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
