package internal

import (
	"errors"
	"fmt"
	"strconv"
)

type PrettyPrinter struct {
	nbCols  int
	headers []string
	data    []map[string]string
}

// Make a new pretty printer with the specified number of cols
func NewPrettyPrinter(headers []string) *PrettyPrinter {
	return &PrettyPrinter{headers: headers, nbCols: len(headers)}
}

// Add a row to the data of the pretty printer
func (pPrinter *PrettyPrinter) AddRow(row map[string]string) error {
	if len(row) == pPrinter.nbCols {
		pPrinter.data = append(pPrinter.data, row)
		return nil
	} else {
		return errors.New("got invalid sized row")
	}
}

func (pPrinter *PrettyPrinter) Print() {

	widths := pPrinter.getWidths()

	fmt.Println(*widths)

	pPrinter.printHeaders(widths)

	for _, row := range pPrinter.data {

		pPrinter.printRow(&row, widths)
	}
}

// Helper function
func (pPrinter *PrettyPrinter) getWidths() *map[string]int {
	max := map[string]int{}

	currentMax := 0

	for _, column := range pPrinter.headers {

		currentMax = len(column)

		for _, row := range pPrinter.data {

			if l := len(row[column]); l < currentMax {
				currentMax = l
			}

		}
		max[column] = currentMax
	}
	return &max
}

func (pPrinter *PrettyPrinter) printHeaders(widths *map[string]int) {
	resultingRow := ""
	prefix := " "
	suffix := " |"

	for _, h := range pPrinter.headers {

		fmt.Println(h)

		format := "%-" + strconv.Itoa((*widths)[h]) + "s"
		content := fmt.Sprintf(format, h)

		resultingRow += prefix + content + suffix
	}

	fmt.Println(resultingRow)

}

// Helper function
func (pPrinter *PrettyPrinter) printRow(row *map[string]string, widths *map[string]int) {

	resultingRow := ""
	prefix := " "
	suffix := " |"

	for _, h := range pPrinter.headers {

		format := "%-" + strconv.Itoa((*widths)[h]) + "s"
		a := (*row)[h]
		content := fmt.Sprintf(format, a)

		resultingRow += prefix + content + suffix
	}

	fmt.Println(resultingRow)
}
