package main

import (
	"os"

	"github.com/taylormonacelli/muchfight"
)

func main() {
	code := muchfight.Execute()
	os.Exit(code)
}
