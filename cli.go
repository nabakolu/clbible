package main

import (
	"os"
)

func cli() {
	book := os.Args[1]
	startVerse := os.Args[2]
	endVerse := os.Args[3]

	parseVerse(book, startVerse, endVerse)
}
