gotable
=======

Go helper to print a table of data to stdout.

## Installation

Make sure you have a working Go environment. Follow the [Go install instructions](http://golang.org/doc/install.html).

To install gotable, simply run:

    go get github.com/muesli/gotable

To compile it from source:

    cd $GOPATH/src/github.com/muesli/gotable
    go get -u -v
    go build

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

## Development

API docs can be found [here](http://godoc.org/github.com/muesli/gotable).

[![Build Status](https://travis-ci.org/muesli/gotable.svg?branch=master)](https://travis-ci.org/muesli/gotable)
