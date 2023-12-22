package main

import (
	"github.com/common-nighthawk/go-figure"
	"os"
)

func main() {

	myFigure := figure.NewFigure(os.Args[1], "", true)
	myFigure.Print()
}
