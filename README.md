# gotable

[![Build Status](https://github.com/muesli/gotable/workflows/build/badge.svg)](https://github.com/muesli/gotable/actions)
[![Go ReportCard](http://goreportcard.com/badge/muesli/gotable)](http://goreportcard.com/report/muesli/gotable)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/muesli/gotable)

Go helper to print a table of data to stdout or an io.Writer.

## Installation

```bash
go get github.com/muesli/gotable
```

## Example

```go
package main

import (
	"github.com/muesli/gotable"
)

func main() {
	tab := gotable.NewTable([]string{"Fruit", "Yumminess", "Color"},
		[]int64{-40, -20, 15},
		"No data in table.")

	tab.AppendRow([]interface{}{"Banana", "very", "yellow"})
	tab.AppendRow([]interface{}{"Pear", "much", "depends"})
	tab.AppendRow([]interface{}{"Pineapple", "delicious", "pineappley"})

	tab.Print()
}
```

## What it looks like
```
Fruit                                     Yumminess                       Color
-------------------------------------------------------------------------------
Banana                                    very                           yellow
Pear                                      much                          depends
Pineapple                                 delicious                  pineappley
```
