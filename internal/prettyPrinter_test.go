package internal

import (
	"math/rand"
	"testing"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyz")

// Generate a random string of length n
func randomString(seed int64, n int) string {
	rand.Seed(seed)

	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func TestPrettyPrinter_AddRow_too_many_column(t *testing.T) {
	var seed int64 = 331

	header1 := randomString(seed, 2)
	header2 := randomString(seed, 3)
	header3 := randomString(seed, 4)

	headers := []string{header1, header2, header3}

	pPrinter := NewPrettyPrinter(headers, "", "")

	row := map[string]string{header1: "col1", header2: "col2", header3: "col3", "added column": "too much"}

	if err := pPrinter.AddRow(row); err == nil {
		t.Errorf("Added a row with too many columns, did not throw an error")
	}

}

func TestPrettyPrinter_getWidths_only_headers(t *testing.T) {

	var seed int64 = 331

	header1 := randomString(seed, 2)
	header2 := randomString(seed, 3)
	header3 := randomString(seed, 4)

	headers := []string{header1, header2, header3}

	pPrinter := NewPrettyPrinter(headers, "", "")

	widths := pPrinter.getWidths()

	for _, header := range headers {

		if got, ok := (*widths)[header]; ok {

			if actual := len(header); actual != got {
				t.Errorf("widths[%s] = %d; want %d", header, got, len(header))
			}

		}
	}

}

func TestPrettyPrinter_getWidhts_first_row_has_bigger_entries(t *testing.T) {

	var seed int64 = 331

	header_lenghts := []int{1, 2, 3, 4}
	headers := make([]string, len(header_lenghts))

	for i := 0; i < len(header_lenghts); i++ {
		headers[i] = randomString(seed, header_lenghts[i])
	}

	row_lengths := []int{3, 4, 5, 6}
	row := make(map[string]string, len(row_lengths))
	for i := 0; i < len(row_lengths); i++ {
		row[headers[i]] = randomString(seed, row_lengths[i])
	}

	pPrinter := NewPrettyPrinter(headers, "", "")
	pPrinter.AddRow(row)

	widths := pPrinter.getWidths()

	for _, header := range headers {

		if got, ok := (*widths)[header]; ok {

			if r, ok := row[header]; ok {

				if actual := len(r); actual != got {
					t.Errorf("widths[%s] = %d; want %d", header, got, actual)

				}
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

func ExamplePrettyPrinter_PrintLastNRows() {
	header1 := "col1"
	header2 := "col2"
	header3 := "col3"

	headers := []string{header1, header2, header3}
	row1 := map[string]string{header1: "col1", header2: "col2", header3: "col3"}
	row2 := map[string]string{header1: "last", header2: "last", header3: "last"}

	pPrinter := NewPrettyPrinter(headers, "", "|")

	pPrinter.AddRow(row1)
	pPrinter.AddRow(row2)

	pPrinter.PrintLastNRows(1)
	// Output:
	// last|last|last

}

func TestPrettyPrinter_PrintLastNRows_n_larger_than_rows(t *testing.T) {
	header1 := "col1"
	header2 := "col2"
	header3 := "col3"

	headers := []string{header1, header2, header3}
	row1 := map[string]string{header1: "col1", header2: "col2", header3: "col3"}
	row2 := map[string]string{header1: "col1", header2: "col2", header3: "col3"}

	pPrinter := NewPrettyPrinter(headers, "", "|")

	pPrinter.AddRow(row1)
	pPrinter.AddRow(row2)

	if err := pPrinter.PrintLastNRows(3); err == nil {
		t.Errorf("Added a row with too many columns, did not throw an error")
	}
}
