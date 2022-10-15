package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	ap "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	red        = "red"
	notRed     = "white"
	threshold  = 32.0
	upperLimit = 35.0
	lowerLimit = 25.0
	inc        = 0.2
	sinkInc    = 5
)

type sensor struct {
	air      string
	sink     string
	power    string
	realSink float64
}

var temperature float64 = 24.0
var up bool = true

func main() {
	a := ap.New()
	w := a.NewWindow("Temperature and Power")

	row1Col1 := widget.NewLabel("Amp")
	row1Col2 := widget.NewLabel("Heatsink")
	row1Col3 := widget.NewLabel("Air")
	row1 := container.New(layout.NewGridLayout(3), row1Col1, row1Col2, row1Col3)

	s := updateSensors()

	row2Col1 := widget.NewLabel(s.power)
	row2Col2 := widget.NewLabel(s.sink)
	row2Col3 := widget.NewLabel(s.air)
	row2 := container.New(layout.NewGridLayout(3), row2Col1, row2Col2, row2Col3)

	grid := container.New(layout.NewVBoxLayout(), row1, row2)
	w.SetContent(grid)
	w.Resize(fyne.NewSize(180, 80))

	go func() {
		colorCode := notRed
		for range time.Tick(time.Second) {
			updateTemp()
			s = updateSensors()
			row2Col1.SetText(s.power)
			row2Col2.SetText(s.sink)
			row2Col3.SetText(s.air)

			if s.realSink > threshold && colorCode != red {
				colorCode = red
				bgColor := canvas.NewRectangle(color.NRGBA{R: 255, G: 0, B: 0, A: 150})
				newRow1 := container.New(layout.NewMaxLayout(), bgColor, row1)
				newRow2 := container.New(layout.NewMaxLayout(), bgColor, row2)
				grid := container.New(layout.NewVBoxLayout(), newRow1, newRow2)
				w.SetContent(grid)
				w.Show()
			}
			if s.realSink < threshold && colorCode == red {
				colorCode = notRed
				grid := container.New(layout.NewVBoxLayout(), row1, row2)
				w.SetContent(grid)
				w.Show()
			}
		}
	}()
	w.Show()
	a.Run()
	fmt.Println("Exited")
}

func updateSensors() *sensor {
	return &sensor{
		power:    "0.0",
		air:      fmt.Sprintf("%0.2f", temperature),
		sink:     fmt.Sprintf("%0.2f", temperature+5),
		realSink: temperature + sinkInc,
	}
}

func updateTemp() {
	if temperature >= upperLimit {
		temperature -= inc
		up = false
	}
	if temperature <= lowerLimit {
		temperature += inc
		up = true
	}
	if temperature > lowerLimit && up {
		temperature += inc
	}
	if temperature < upperLimit && !up {
		temperature -= inc
	}
}
