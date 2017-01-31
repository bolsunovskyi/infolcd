package main

import (
	ui "github.com/gizak/termui"
	"time"
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
	col := ui.NewCol(4, 4, timePar)

	ui.Body.AddRows(ui.NewRow(col));
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
