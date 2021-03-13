package internal

import (
	"errors"
	"fmt"
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

	prettyPrinter.PrintRow(prettyPrinter.header)

	for _, row := range prettyPrinter.data {

		prettyPrinter.PrintRow(row)
	}
}

func (prettyPrinter *PrettyPrinter) PrintRow(row []string) {

	resultingRow := ""

	for _, s := range row {
		resultingRow += s + "     |"
	}

	fmt.Println(resultingRow)
}
