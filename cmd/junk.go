package main

//
// import (
// 	"fmt"
// 	"image/color"
// 	"time"
//
// 	"fyne.io/fyne/v2"
// 	ap "fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/layout"
// 	"fyne.io/fyne/v2/widget"
// )
//
// const (
// 	red    = "red"
// 	notRed = "white"
// )
//
// func main() {
// 	a := ap.New()
// 	w := a.NewWindow("Temperature and Power")
//
// 	row1 := []string{"Amp", "Heatsink", "Air"}
// 	row3 := []string{"0.0", "27.0", "28.5"}
//
// 	grid := getGrid(row1, row3, notRed)
//
// 	t := 25.0
//
// 	w.SetContent(grid)
// 	w.Resize(fyne.NewSize(180, 80))
// 	go func() {
// 		var grid *fyne.Container
// 		level := 1 //white background incrementing
// 		for range time.Tick(time.Second) {
// 			switch {
// 			case t > 35 && level == 3:
// 				level = 4
// 				t -= 0.2
// 			case t > 35 && level == 4:
// 				level = 4
// 				t -= 0.2
// 			case t > 30 && level == 2:
// 				level = 3
// 				t += 0.2
// 			case t > 30 && level == 3:
// 				level = 3
// 				t += 0.2
// 			case t > 30 && level == 4:
// 				level = 5
// 				t -= 0.2
// 			case t > 30 && level == 5:
// 				level = 5
// 				t -= 0.2
// 			case t > 25 && level == 1:
// 				level = 2
// 				t += 0.2
// 			case t > 25 && level == 2:
// 				level = 2
// 				t += 0.2
// 			case t > 25 && level == 5:
// 				level = 6
// 				t -= 0.2
// 			case t > 25 && level == 6:
// 				level = 6
// 				t -= 0.2
// 			default:
// 				level = 1
// 				t += 0.2
// 			}
// 			row3[2] = fmt.Sprintf("%0.1f", t)
// 			if level == 3 || level == 5 || level == 4 {
// 				grid = getGrid(row1, row3, red)
// 			} else {
// 				grid = getGrid(row1, row3, notRed)
// 			}
// 			w.SetContent(grid)
// 			w.Show()
// 		}
// 	}()
// 	w.Show()
// 	a.Run()
// 	fmt.Println("Exited")
// }
//
// func rowData(s []string, colorCode string) *fyne.Container {
// 	c := container.New(layout.NewGridLayout(3),
// 		widget.NewLabel(s[0]),
// 		widget.NewLabel(s[1]),
// 		widget.NewLabel(s[2]))
//
// 	if colorCode == "red" {
// 		bgColor := canvas.NewRectangle(color.NRGBA{R: 255, G: 0, B: 0, A: 150})
// 		return container.New(layout.NewMaxLayout(), bgColor, c)
// 	}
// 	return c
// }
//
// func getGrid(row1, row3 []string, colorCode string) *fyne.Container {
// 	return container.New(layout.NewVBoxLayout(),
// 		rowData(row1, colorCode),
// 		rowData(row3, colorCode))
// }
//
