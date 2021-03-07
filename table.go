package gotable

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"
)

// Table is a helper for printing data in sheet form
type Table struct {
	headers   []string
	widths    []int64
	rows      [][]interface{}
	summary   []interface{}
	emptyText string
	writer    io.Writer
}

// NewTable returns a new table
func NewTable(headers []string, widths []int64, emptyText string) Table {
	return NewTableWithWriter(headers, widths, emptyText, os.Stdout)
}

// NewTable returns a new table
func NewTableWithWriter(headers []string, widths []int64, emptyText string, writer io.Writer) Table {
	return Table{
		headers:   headers,
		widths:    widths,
		rows:      [][]interface{}{},
		emptyText: emptyText,
		writer:    writer,
	}
}

func (t *Table) checkWidths(row []interface{}) {
	for col, v := range row {
		s := fmt.Sprint(v)
		sl := runewidth.StringWidth(s)
		if float64(sl) > math.Abs(float64(t.widths[col])) {
			// the added value doesn't fit in this column, let's make it bigger
			l := int64(sl)
			if t.widths[col] < 0 {
				l *= -1
			}
			t.widths[col] = l
		}
	}
}

// AppendRow adds a row to the end of the table
func (t *Table) AppendRow(row []interface{}) {
	t.checkWidths(row)
	t.rows = append(t.rows, row)
}

// SetSummary sets a summary for this table, to be printed below the table
func (t *Table) SetSummary(summary []interface{}) {
	t.checkWidths(summary)
	t.summary = summary
}

// Print writes the entire table to stdout
func (t *Table) Print() error {
	totalWidth := int64(0)
	format := ""
	for i, w := range t.widths {
		margin := 2
		if i == len(t.widths)-1 {
			// Don't add margin for the last column
			margin = 0
		}
		format += "%" + strconv.FormatInt(w, 10) + "s" + strings.Repeat(" ", margin)
		totalWidth += int64(math.Abs(float64(w))) + int64(margin)
	}

	// print header
	if _, err := fmt.Fprintf(t.writer, format+"\n", ifaceify(t.headers)...); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(t.writer, strings.Repeat("-", int(totalWidth))+"\n"); err != nil {
		return err
	}

	// print rows
	for _, row := range t.rows {
		if _, err := fmt.Fprintf(t.writer, format+"\n", row...); err != nil {
			return err
		}
	}
	if len(t.rows) == 0 {
		if _, err := fmt.Fprintf(t.writer, t.emptyText+"\n"); err != nil {
			return err
		}
	} else if len(t.summary) > 0 {
		t.PrintSummary()
	}

	return nil
}

// PrintSummary writes the table summary to stdout
func (t *Table) PrintSummary() error {
	totalWidth := int64(0)
	format := ""
	for i, w := range t.widths {
		margin := 2
		if i == len(t.widths)-1 {
			// Don't add margin for the last column
			margin = 0
		}
		format += "%" + strconv.FormatInt(w, 10) + "s" + strings.Repeat(" ", margin)
		totalWidth += int64(math.Abs(float64(w))) + int64(margin)
	}

	// print divider
	if _, err := fmt.Fprintf(t.writer, strings.Repeat("-", int(totalWidth))+"\n"); err != nil {
		return err
	}

	// print summary
	if _, err := fmt.Fprintf(t.writer, format+"\n", t.summary...); err != nil {
		return err
	}

	return nil
}

func ifaceify(list []string) []interface{} {
	vals := make([]interface{}, len(list))
	for i, v := range list {
		vals[i] = v
	}
	return vals
}
