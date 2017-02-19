package main

import (
	"github.com/bolsunovskyi/infolcd/gui"
	"github.com/bolsunovskyi/infolcd/exchange"
	"github.com/bolsunovskyi/infolcd/weather"

	"flag"
	"log"
	"fmt"
)

var cityID int
var weatherApiKey string

type wt struct {}

func (w wt) Update(r *weather.Response) {
	gui.Weather = fmt.Sprintf("%0.1f, %s\n%s", r.Main.Temp, r.Weather[0].Main, r.Weather[0].Description)
	gui.Update()
}

type ex struct {}

func (e ex) Update(r *exchange.Response) {
	gui.Exchange = fmt.Sprintf("%0.2f/%0.2f", r.LatestRates.WholeBuy, r.LatestRates.WholeSale)
	gui.Update()
}

func init() {
	flag.IntVar(&cityID, "c", 703448, "City ID")
	flag.StringVar(&weatherApiKey, "k", "", "Open Weather Map api key")
	flag.Parse()
}

func main() {
	if weatherApiKey == "" {
		log.Println("Open Weather api key is required")
		return
	}

	gui.Init()
	defer gui.Close()

	weather.Listen(wt{}, cityID, weatherApiKey)
	exchange.Listen(ex{})

	gui.Loop()
}
