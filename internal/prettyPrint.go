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
	prefix  string
	suffix  string
}

// Make a new pretty printer with the specified number of cols
func NewPrettyPrinter(headers []string) *PrettyPrinter {
	return &PrettyPrinter{headers: headers, nbCols: len(headers), prefix: " ", suffix: " |"}
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

// func (pPrinter *PrettyPrinter) Print() {

// 	widths := pPrinter.getWidths()

// 	fmt.Println(*widths)

// 	pPrinter.printHeaders(widths)

// 	for _, row := range pPrinter.data {

// 		pPrinter.printRow(&row, widths)
// 	}
// }

// Helper function
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

func (pPrinter *PrettyPrinter) PrintHeaders() {
	resultingRow := ""
	widths := pPrinter.getWidths()

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

// Helper function
func (pPrinter *PrettyPrinter) PrintRow(row map[string]string) {

	resultingRow := ""
	widths := pPrinter.getWidths()

	for i, h := range pPrinter.headers {

		format := "%-" + strconv.Itoa((*widths)[h]) + "s"
		a := row[h]
		content := fmt.Sprintf(format, a)

		if i == len(pPrinter.headers)-1 {
			resultingRow += pPrinter.prefix + content
		} else {
			resultingRow += pPrinter.prefix + content + pPrinter.suffix
		}
	}

	fmt.Println(resultingRow)
}
