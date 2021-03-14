package internal

import (
	"math/rand"
	"testing"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyz")

func randomString(seed int64, n int) string {
	rand.Seed(seed)

	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func TestWidhts(t *testing.T) {

	var seed int64 = 331

	header1 := randomString(seed, 2)
	header2 := randomString(seed, 3)
	header3 := randomString(seed, 4)

	headers := []string{header1, header2, header3}

	pPrinter := NewPrettyPrinter(headers, "", "")

	widths := pPrinter.getWidths()

	for _, header := range headers {

		if got, ok := (*widths)[header]; ok {

			if len(header) != got {
				t.Errorf("widths[%s] = %d; want %d", header, got, len(header))
			}

		}
	}

}

func ExamplePrettyPrinter_printHeaders_threeHeaders() {

	header1 := "Dit is header 1"
	header2 := "Dit is header 2"
	header3 := "Dit is header 3"

	headers := []string{header1, header2, header3}
	pPrinter := NewPrettyPrinter(headers, "", "|")

	widths := pPrinter.getWidths()

	pPrinter.printHeaders(widths)
	// Output:
	// Dit is header 1|Dit is header 2|Dit is header 3

}

func ExamplePrettyPrinter_printRow() {

	header1 := "col1 "
	header2 := "col2  "
	header3 := "col3  "

	headers := []string{header1, header2, header3}
	row := map[string]string{header1: "col1", header2: "col2", header3: "col3"}

	pPrinter := NewPrettyPrinter(headers, "", "|")
	widths := pPrinter.getWidths()

	pPrinter.AddRow(row)

	pPrinter.printRow(&row, widths)
	// Output:
	// col1 |col2  |col3
}

func ExamplePrettyPrinter_Print() {

	header1 := "col1"
	header2 := "col22"
	header3 := "col333"

	headers := []string{header1, header2, header3}
	row := map[string]string{header1: "col1", header2: "col2", header3: "col3"}

	pPrinter := NewPrettyPrinter(headers, "", "|")

	pPrinter.AddRow(row)

	pPrinter.Print()
	// Output:
	//col1|col22|col333
	//col1|col2 |col3
}

func ExamplePrettyPrinter_Print_wrong_entry_in_row() {
	header1 := "col1"
	header2 := "12345"
	header3 := "col333"

	headers := []string{header1, header2, header3}
	row := map[string]string{header1: "col1", "wrong_header": "col2", header3: "col3"}

	pPrinter := NewPrettyPrinter(headers, "", "|")

	pPrinter.AddRow(row)

	pPrinter.Print()
	// Output:
	// col1|12345|col333
	// col1|     |col3

}
