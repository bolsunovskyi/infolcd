package main

import (
	ui "github.com/gizak/termui"
	"time"
	"github.com/bolsunovskyi/infolcd/exchange"
	"fmt"
	"github.com/bolsunovskyi/infolcd/weather"
)

func main() {



	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	timePar := ui.NewPar(time.Now().Format("02.01.2006\n15:04:05"))
	timePar.Height = 4
	timePar.BorderLabel = "Date/Time"
	timePar.Width = 5

	usd, err := exchange.GetUSD()
	var usdVal string
	if err != nil {
		usdVal = "Error :("
	} else {
		usdVal = fmt.Sprintf("%.2f/%.2f", usd.LatestRates.WholeBuy, usd.LatestRates.WholeSale)
	}

	excPar := ui.NewPar(usdVal)
	excPar.Height = 4
	excPar.BorderLabel = "Exchange"
	excPar.Width = 5

	w, err := weather.GetWeather()
	var weatherVal string
	if err != nil {
		weatherVal = "Error :("
	} else {
		weatherVal = fmt.Sprintf("%.1f, %s\n%s", w.Main.Temp, w.Weather[0].Main, w.Weather[0].Description)
	}

	weatherPar := ui.NewPar(weatherVal)
	weatherPar.Height = 4
	weatherPar.BorderLabel = "Weather"
	weatherPar.Width = 5


	excCol := ui.NewCol(6, 0, excPar)
	timeCol := ui.NewCol(6, 0, timePar)
	weatherCol := ui.NewCol(6, 0, weatherPar)


	ui.Body.AddRows(ui.NewRow(timeCol, excCol));
	ui.Body.AddRows(ui.NewRow(weatherCol));
	ui.Body.Align()
	ui.Render(ui.Body)


	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Handle("/timer/1s", func(e ui.Event) {
		timePar.Text = time.Now().Format("02.01.2006\n15:04:05");
		ui.Render(ui.Body)
	})



	ui.Loop()
}
