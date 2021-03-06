package gui

import (
	"github.com/gizak/termui"
	"time"
)

const parHeight = 4
const parWidth = 5

var Exchange string
var Weather string
var Temp string

var timePar, exchangePar, weatherPar, tempPar *termui.Par

func Init() {
	Exchange = "0/0"
	Weather = "0"
	Temp = "0"

	timePar = termui.NewPar(time.Now().Format("02.01.2006\n15:04:05"))
	timePar.Width = parWidth
	timePar.Height = parHeight
	timePar.BorderLabel = "Дата/Время"

	exchangePar = termui.NewPar(Exchange)
	exchangePar.Width = parWidth
	exchangePar.Height = parHeight
	exchangePar.BorderLabel = "Курс $"

	weatherPar = termui.NewPar(Weather)
	weatherPar.Width = parWidth
	weatherPar.Height = parHeight
	weatherPar.BorderLabel = "Погода"

	tempPar = termui.NewPar(Temp)
	tempPar.Width = parWidth
	tempPar.Height = parHeight
	tempPar.BorderLabel = "Температура"

	err := termui.Init()
	if err != nil {
		panic(err)
	}

	termui.Body.AddRows(termui.NewRow(termui.NewCol(6, 0, timePar), termui.NewCol(6, 0, exchangePar)))
	termui.Body.AddRows(termui.NewRow(termui.NewCol(6, 0, weatherPar), termui.NewCol(6, 0, tempPar)))
	termui.Body.Align()
	termui.Render(termui.Body)

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		termui.Body.Width = termui.TermWidth()
		termui.Body.Align()
		termui.Clear()
		termui.Render(termui.Body)
	})

	termui.Handle("/timer/1s", func(e termui.Event) {
		timePar.Text = time.Now().Format("02.01.2006\n15:04:05")
		termui.Render(termui.Body)
	})
}

func Loop() {
	termui.Loop()
}

func Close() {
	termui.Close()
}

func Update() {
	timePar.Text = time.Now().Format("02.01.2006\n15:04:05")
	exchangePar.Text = Exchange
	weatherPar.Text = Weather
	tempPar.Text = Temp

	termui.Render(termui.Body)
}
