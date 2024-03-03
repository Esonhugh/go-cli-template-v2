package Print

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// Table is Struct for printable table
type Table struct {
	// Table Header as []string{"id","name","value" .....}
	Header []string
	// Table Body is content of table.
	Body [][]string

	ShrinkBigTableInPrint bool               // ShrinkBigTable will shrink the table if it's too large.
	ShrinkBigTableFunc    func(table *Table) // ShrinkTableFunc is a function to shrink the table.
}

// DefaultShrinkTableFunc is the default function to shrink the table.
// It will shrink the table to 7 columns and width of each unit is less than 50 characters.
func DefaultShrinkTableFunc(table *Table) {
	if len(table.Header) > 7 {
		table.Header = table.Header[:6]
		for i, v := range table.Body {
			table.Body[i] = v[:6]
		}
	}
	for rowID, row := range table.Body {
		for unitID, unit := range row {
			lines := strings.Split(unit, "\n")
			for i, line := range lines {
				if len(line) > 50 {
					lines[i] = line[:47] + "..."
				}
			}
			table.Body[rowID][unitID] = strings.Join(lines, "\n")
		}
	}
	return
}

// Print function prints the table itself.
func (t *Table) Print(Caption string) {
	// If Table is too large. Hacker can use export CFP_TABLE_LOG=true to print it to file cfp_debug.log
	if t.ShrinkBigTableInPrint {
		if t.ShrinkBigTableFunc != nil {
			t.ShrinkBigTableFunc(t)
		}
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(t.Header)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	var TableHeaderColor = make([]tablewriter.Colors, len(t.Header))
	for i := range TableHeaderColor {
		TableHeaderColor[i] = tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor}
	}
	table.SetHeaderColor(TableHeaderColor...)
	if Caption != "" {
		table.SetCaption(true, Caption)
	}
	table.AppendBulk(t.Body)
	table.Render()
}
