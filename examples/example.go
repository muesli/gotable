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
