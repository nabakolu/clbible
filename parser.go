package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseVerse(book string, startVerse string, endVerse string) {
	file, err := os.Open(Config.TranslationsDir + "/" + Config.Translation + "/" + book + ".usx")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	correctSid := book + " " + startVerse
	correctEid := book + " " + endVerse
	inCorrectVerses := false
	inNote := false
	noteCounter := 0
	initialSpace := true
	var notes []string

	decoder := xml.NewDecoder(file)

parsing:
	for {
		tok, err := decoder.Token()
		if err != nil {
			break
		}

		switch v := tok.(type) {
		case xml.StartElement:
			switch v.Name.Local {
			case "verse":
				verseNumber := ""
				startOfVerse := false
				endOfVerse := false
				for _, attr := range v.Attr {
					if attr.Name.Local == "sid" {
						startOfVerse = true
						if attr.Value == correctSid {
							inCorrectVerses = true
						}
					}
					if attr.Name.Local == "eid" {
						endOfVerse = true
						if attr.Value == correctEid {
							break parsing
						}
					}
					if attr.Name.Local == "number" {
						verseNumber = attr.Value
					}
				}
				if startOfVerse && inCorrectVerses && Config.ShowVerseNumbers {
					fmt.Print(subScriptNumber(verseNumber))
					initialSpace = true
				}
				if endOfVerse && inCorrectVerses {
					fmt.Print(" ")
				}
			case "note":
				if inCorrectVerses {
					inNote = true
					if Config.ShowNotes {
						fmt.Print(superScriptNumber(noteCounter + 1))
						notes = append(notes, "")
					}
				}
			case "chapter":
				if inCorrectVerses {
					for _, attr := range v.Attr {
						if attr.Name.Local == "sid" {
							fmt.Println("\n\n" + attr.Value)
						}
					}
				}
			}
		case xml.EndElement:
			switch v.Name.Local {
			case "note":
				if inCorrectVerses {
					inNote = false
					noteCounter++
				}
			}
		case xml.CharData:
			text := strings.TrimSpace(string(v))
			if text == "" {
				continue
			}
			if !initialSpace && text[0] != ',' && text[0] != '.' && text[0] != '?' && text[0] != '!' && text[0] != ';' {
				text = " " + text
			}
			if inCorrectVerses && !inNote {
				initialSpace = false
				fmt.Print(text)
			} else if inCorrectVerses && inNote && Config.ShowNotes {
				notes[noteCounter] += text
			}
		}
	}
	fmt.Println()
	if Config.ShowNotes {
		fmt.Println()
		for i := range notes {
			fmt.Println(superScriptNumber(i+1) + notes[i])
		}
	}
}

func superScriptNumber(number int) string {
	superScriptNumber := map[rune]rune{
		'0': '⁰', '1': '¹', '2': '²', '3': '³', '4': '⁴', '5': '⁵', '6': '⁶', '7': '⁷', '8': '⁸', '9': '⁹',
	}
	numberString := strconv.Itoa(number)
	var result string
	for _, char := range numberString {
		result += string(superScriptNumber[char])
	}
	return result
}

func subScriptNumber(numberString string) string {
	subScriptNumber := map[rune]rune{
		'0': '₀', '1': '₁', '2': '₂', '3': '₃', '4': '₄', '5': '₅', '6': '₆', '7': '₇', '8': '₈', '9': '₉',
	}
	var result string
	for _, char := range numberString {
		result += string(subScriptNumber[char])
	}
	return result
}
