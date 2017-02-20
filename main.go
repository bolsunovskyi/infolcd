package main

import (
	ex "github.com/bolsunovskyi/infolcd/exchange"
	wt "github.com/bolsunovskyi/infolcd/weather"
	"github.com/bolsunovskyi/pb_api"
	"github.com/bolsunovskyi/infolcd/gui"

	"flag"
	"log"
	"fmt"
)

type weather struct {
	CityID	int
	APIKey	string

}

type exchange struct {
	PBClient	string
	PBSecret	string
}

func (w weather) Update(r *wt.Response) {
	gui.Weather = fmt.Sprintf("%0.1f, %s\n%s", r.Main.Temp, r.Weather[0].Main, r.Weather[0].Description)
	gui.Update()
}

func (w weather) GetAPIKey () string {
	return w.APIKey
}

func (w weather) GetCityID () int {
	return w.CityID
}

func (e exchange) Update(r *pb_api.ExchangeRate) {
	gui.Exchange = fmt.Sprintf("%0.2f/%0.2f", r.GetBuy(), r.GetSale())
	gui.Update()
}

func (e exchange) GetPBID () string {
	return e.PBClient
}

func (e exchange) GetPBSecret () string {
	return e.PBSecret
}

var _weather weather
var _exchange exchange

func init() {
	flag.IntVar(&_weather.CityID, "c", 703448, "City ID")
	flag.StringVar(&_weather.APIKey, "k", "", "Open Weather Map api key")
	flag.StringVar(&_exchange.PBClient, "pbid", "", "Privat bank api client id")
	flag.StringVar(&_exchange.PBSecret, "pbsecret", "", "Privat bank api client id")
	flag.Parse()
}

func main() {
	if _weather.APIKey == "" {
		log.Println("Open Weather api key is required")
		return
	}

	if _exchange.PBSecret == "" {
		log.Println("PB secret is required")
		return
	}

	if _exchange.PBClient == "" {
		log.Println("PB client id is requird")
		return
	}

	gui.Init()
	defer gui.Close()

	wt.Listen(_weather)
	ex.Listen(_exchange)

	gui.Loop()
}
