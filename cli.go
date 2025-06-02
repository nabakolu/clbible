package main

import (
	"flag"
	"fmt"
)

func cli() {
	var version bool
	var yesNotes bool
	var noNotes bool
	var yesNums bool
	var noNums bool

	flag.StringVar(&Config.Translation, "translation", Config.Translation, "Trannslation used")
	flag.StringVar(&Config.TranslationsDir, "translationsDir", Config.TranslationsDir, "Directory, where the translations reside in")
	flag.BoolVar(&version, "version", false, "Print version")
	flag.BoolVar(&yesNotes, "yesNotes", false, "Show notes")
	flag.BoolVar(&noNotes, "noNotes", false, "Don't show notes")
	flag.BoolVar(&yesNums, "yesNums", false, "Show verse numbers")
	flag.BoolVar(&noNums, "noNums", false, "Don't show verse numbers")

	flag.Parse()

	if yesNotes {
		Config.ShowNotes = true
	}
	if noNotes {
		Config.ShowNotes = false
	}
	if yesNums {
		Config.ShowVerseNumbers = true
	}
	if noNums {
		Config.ShowVerseNumbers = false
	}


	if version {
		fmt.Println(Version)
		return
	}

	args := flag.Args()

	if len(args) < 2 {
		flag.Usage()
		return
	}

	book := args[0]
	startVerse := args[1]
	var endVerse string
	if len(args) > 2 {
		endVerse = args[2]
	} else {
		endVerse = startVerse
	}

	parseVerse(book, startVerse, endVerse)
}
