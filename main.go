package main

import (
	ui "github.com/gizak/termui"
	"time"
	"github.com/bolsunovskyi/infolcd/exchange"
	"fmt"
)

func main() {



	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	timePar := ui.NewPar(time.Now().Format("02.01.2006 15:04:05"))
	timePar.Height = 3
	timePar.BorderLabel = "Date/Time"
	timePar.Width = 5

	usd, err := exchange.GetUSD()
	var usdVal string
	if err != nil {
		usdVal = err.Error()
	} else {
		usdVal = fmt.Sprintf("%.2f/%.2f", usd.LatestRates.WholeBuy, usd.LatestRates.WholeSale)
	}

	excPar := ui.NewPar(usdVal)
	excPar.Height = 3
	excPar.BorderLabel = "Exchange Rate"
	excPar.Width = 5


	excCol := ui.NewCol(4, 0, excPar)
	timeCol := ui.NewCol(4, 0, timePar)


	ui.Body.AddRows(ui.NewRow(timeCol, excCol));
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
		timePar.Text = time.Now().Format("02.01.2006 15:04:05");
		ui.Render(ui.Body)
	})



	ui.Loop()
}
