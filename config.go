package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type config struct {
	TranslationsDir  string `json:"translationsDir,omitempty"`
	Translation      string `json:"translation,omitempty"`
	ShowVerseNumbers bool   `json:"showVerseNumbers,omitempty"`
	ShowNotes        bool   `json:"showNotes,omitempty"`
}

var Config config

var Version = "0.2"

func readConfig() {
	Config.Translation = "ELB"
	Config.ShowVerseNumbers = true
	Config.ShowNotes = false
	// Get the home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	Config.TranslationsDir = filepath.Join(homeDir, ".clbible", "translations")

	// Path to the config file
	configFilePath := filepath.Join(homeDir, ".clbible", "config.json")

	// Open the config file
	file, err := os.Open(configFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	// Parse the JSON config into the Config variable
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Config); err != nil {
		fmt.Println("unable to parse config file: %w", err)
	}

	// Expand ~, $HOME, or other env variables
	Config.TranslationsDir = os.ExpandEnv(Config.TranslationsDir)
}
