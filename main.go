package main

import (
	"github.com/bolsunovskyi/infolcd/gui"
	"flag"
	"log"
)

var cityID int
var weatherApiKey string

func init() {
	flag.IntVar(&cityID, "c", 703448, "City ID")
	flag.StringVar(&weatherApiKey, "k", "", "Open Weather Map api key")
	flag.Parse()

	if weatherApiKey == "" {
		log.Fatalln("Open Weather api key is required")
	}
}

func main() {
	defer gui.Close()
	gui.Loop()
}
