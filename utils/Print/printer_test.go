package Print

import "testing"

func TestTablePrintWithNewlines(t *testing.T) {
	tab := Table{
		Header: []string{"Header1", "Header2"},
		Body: [][]string{
			{"Row1", "Row2"},
		},
	}
	tab.Print("testTablePrintWithNewlines")
	tab.Body = [][]string{
		{"Row1\nRow1", "Row2"},
	}
	tab.Print("testTablePrintWithNewlines")
}
