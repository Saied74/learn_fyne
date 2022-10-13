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

func main() {
	a := ap.New()
	w := a.NewWindow("Temperature and Power")

	// clock := widget.NewLabel("xxx")
	// updateTime(clock)

	txt1 := canvas.NewText("Amp", color.Black)
	txt2 := canvas.NewText("Heatsink", color.Black)
	txt3 := canvas.NewText("Air", color.Black)
	txt4 := canvas.NewText("0.0 W", color.Black)
	txt5 := canvas.NewText("27 degC", color.Black)
	txt6 := canvas.NewText("25 degC", color.Black)
	grid := container.New(layout.NewGridLayout(3), txt1, txt2, txt3, txt4, txt5, txt6)

	w.SetContent(grid)
	w.Resize(fyne.NewSize(100, 50))
	// go func() {
	// 	for range time.Tick(time.Second) {
	// 		updateTime(clock)
	// 	}
	// }()
	w.Show()
	a.Run()
	fmt.Println("Exited")
}

func updateTime(clock *widget.Label) {
	f := time.Now().Format("Time: 03:04:05")
	clock.SetText(f)
}
