package main

type config struct {
	translationsDir  string
	translation      string
	showVerseNumbers bool
	showNotes        bool
}

var Config config

func readConfig() {
	Config.translation = "ELB"
	Config.translationsDir = "./translations/"
	Config.showVerseNumbers = true
	Config.showNotes = false
}
