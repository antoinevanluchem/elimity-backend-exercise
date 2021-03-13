package internal

import (
	"errors"
	"fmt"
	"strconv"
)

type PrettyPrinter struct {
	nbCols int
	header []string
	data   [][]string
}

// Make a new pretty printer with the specified number of cols
func NewPrettyPrinter(nbCols int) *PrettyPrinter {
	return &PrettyPrinter{nbCols: nbCols}
}

// Add a row to the data of the pretty printer
func (prettyPrinter *PrettyPrinter) AddRow(row []string) error {
	if len(row) == prettyPrinter.nbCols {
		prettyPrinter.data = append(prettyPrinter.data, row)
		return nil
	} else {
		return errors.New("got invalid sized row")
	}
}

// Set the headers of the pretty printer
func (prettyPrinter *PrettyPrinter) SetHeader(header []string) error {
	if len(header) == prettyPrinter.nbCols {
		prettyPrinter.header = header
		return nil
	} else {
		return errors.New("got invalid sized headers")
	}
}

func (prettyPrinter *PrettyPrinter) Print() {

	width := prettyPrinter.getMaximalWidth()

	fmt.Printf("Width is %d\n", width)

	prettyPrinter.printRow(prettyPrinter.header, width)

	for _, row := range prettyPrinter.data {

		prettyPrinter.printRow(row, width)
	}
}

// Helper function
func (prettyPrinter *PrettyPrinter) getMaximalWidth() (max int) {

	max = 0
	for _, v := range prettyPrinter.header {

		if l := len(v); l > max {
			max = l
		}
	}

	for _, v := range prettyPrinter.data {

		if l := len(v); l > max {
			max = l
		}
	}

	return
}

// Helper function
func (prettyPrinter *PrettyPrinter) printRow(row []string, width int) {

	resultingRow := ""

	for _, s := range row {
		resultingRow += fmt.Sprintf("%-"+strconv.Itoa(width+2)+"s", s) + "|"
	}

	fmt.Println(resultingRow)
}
