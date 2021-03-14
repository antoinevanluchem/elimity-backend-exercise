package internal

import (
	"errors"
	"fmt"
	"strconv"
)

// A struct that defines the pretty printer
// nbCols 		= number of columns of the table
// headers 		= headers (first row) of the table
// data			= all other rows of the table
// prefix 		= the prefix of the content of an element in the table
// suffix 		= the suffix of the content of an element in the table
type PrettyPrinter struct {
	nbCols  int
	headers []string
	data    []map[string]string
	prefix  string
	suffix  string
}

// Make a new pretty printer
func NewPrettyPrinter(headers []string, prefix string, suffix string) *PrettyPrinter {
	return &PrettyPrinter{headers: headers, nbCols: len(headers), prefix: prefix, suffix: suffix}
}

// Add a row to the data of the pretty printer
func (pPrinter *PrettyPrinter) AddRow(row map[string]string) error {
	if len(row) == pPrinter.nbCols {
		pPrinter.data = append(pPrinter.data, row)
		return nil
	}

	return errors.New("got invalid sized row")

}

// Print the whole table: headers and data
func (pPrinter *PrettyPrinter) Print() {

	widths := pPrinter.getWidths()

	pPrinter.printHeaders(widths)

	for _, row := range pPrinter.data {

		pPrinter.printRow(&row, widths)
	}
}

// Print the last N rows of the data
func (pPrinter *PrettyPrinter) PrintLastNRows(n int) {

	widths := pPrinter.getWidths()

	for _, row := range pPrinter.data[len(pPrinter.data)-n:] {
		pPrinter.printRow(&row, widths)
	}
}

// Helper function to determine the widths of every column
func (pPrinter *PrettyPrinter) getWidths() *map[string]int {
	max := map[string]int{}

	currentMax := 0

	for _, column := range pPrinter.headers {

		currentMax = len(column)

		for _, row := range pPrinter.data {

			if l := len(row[column]); currentMax < l {
				currentMax = l
			}

		}
		max[column] = currentMax
	}
	return &max
}

// Helper function to print the header row
func (pPrinter *PrettyPrinter) printHeaders(widths *map[string]int) {
	resultingRow := ""

	for i, h := range pPrinter.headers {

		format := "%-" + strconv.Itoa((*widths)[h]) + "s"
		content := fmt.Sprintf(format, h)

		if i == len(pPrinter.headers)-1 {
			resultingRow += pPrinter.prefix + content
		} else {
			resultingRow += pPrinter.prefix + content + pPrinter.suffix
		}
	}

	fmt.Println(resultingRow)

}

// Helper function to print a row that is not a header row
func (pPrinter *PrettyPrinter) printRow(row *map[string]string, widths *map[string]int) {

	resultingRow := ""

	for i, h := range pPrinter.headers {

		format := "%-" + strconv.Itoa((*widths)[h]) + "s"

		content := ""
		if a, ok := (*row)[h]; ok {
			content = fmt.Sprintf(format, a)
		}

		if i == len(pPrinter.headers)-1 {
			resultingRow += pPrinter.prefix + content
		} else {
			resultingRow += pPrinter.prefix + content + pPrinter.suffix
		}
	}

	fmt.Println(resultingRow)
}
