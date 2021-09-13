package main

import(
	"os"
	"github.com/common-nighthawk/go-figure"
)

func main() {

  myFigure := figure.NewFigure(os.Args[1],"", true)
  myFigure.Print()
}
