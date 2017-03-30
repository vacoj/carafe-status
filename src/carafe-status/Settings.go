package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

type ISettings interface {
	parseSettings()
}

type Settings struct {
	WebPort    string   `json:"web_port"`
	MOTD       string   `json:"motd"`
	AlertEmail string   `json:"alert_email"`
	Carafes    []Carafe `json:"carafes"`
}

func (s *Settings) load() {
	confFile := "../../init/config.json"
	if len(os.Args) > 1 {
		confFile = os.Args[1]
	}

	file, err := os.Open(confFile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	fileContent, err := os.Open(confFile)
	if err != nil {
		fmt.Println("Could not open config file")
	}

	jsonParser := json.NewDecoder(fileContent)
	if err = jsonParser.Decode(&s); err != nil {
		fmt.Println("Could not load config file. Check JSON formatting.")
	}
	spew.Dump(SETTINGS)
}
