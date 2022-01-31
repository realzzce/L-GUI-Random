package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func random(min, max int) int {
	if min == max {
		max = max + 1
	}
	if min > max {
		temp := max
		max = min
		min = temp
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

func main() {
	var logStr string
	// default 0 - 99
	var Max, Min int
	Min = 0
	Max = 100
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Hi")
	w.Resize(fyne.NewSize(300, 150))

	inputMin := widget.NewEntry()
	inputMin.SetPlaceHolder("MIN ...")
	inputMax := widget.NewEntry()
	inputMax.SetPlaceHolder("MAX - 1 ...")

	inputMin.OnChanged = func(content string) {
		inputA := strings.TrimSpace(inputMin.Text)
		if inputMin.Text == "" || inputMin.Text == " " {
			Min = 0
		} else {
			MinNumber, errA := strconv.Atoi(inputA)
			if errA == nil {
				Min = MinNumber
			} else {
				Min = 0
			}
		}

	}
	inputMax.OnChanged = func(content string) {
		Max = 100
		inputB := strings.TrimSpace(inputMax.Text)
		if inputMax.Text == "" || inputMax.Text == " " {
			Max = 100
		} else {
			MaxNumber, errB := strconv.Atoi(inputB)
			if errB == nil {
				Max = MaxNumber
			} else {
				Max = 100
			}
		}

	}
	hello := widget.NewLabel("Hello")
	logNum := widget.NewLabel("")

	inputContent := container.New(layout.NewGridLayoutWithColumns(2),
		inputMin, inputMax)

	buttonContent := container.New(layout.NewGridLayoutWithColumns(3),
		widget.NewButton("Random", func() {
			randomtext := fmt.Sprint(random(Min, Max))
			hello.SetText(randomtext)
		}),
		widget.NewButton("Random * 10", func() {
			logStr = ""
			for i := 1; i <= 10; i++ {
				randomtext := fmt.Sprint(random(Min, Max))
				hello.SetText(randomtext)
				if i == 10 {
					logStr = logStr + randomtext
				} else {
					logStr = logStr + randomtext + " , "
				}
				logNum.SetText(logStr)
			}
		}),
		widget.NewButton("Clear", func() {
			hello.SetText("")
			logNum.SetText("")
		}),
	)

	// layout
	w.SetContent(container.New(layout.NewVBoxLayout(),
		inputContent,
		widget.NewSeparator(),
		hello,
		widget.NewSeparator(),
		buttonContent,
		logNum))

	w.ShowAndRun()
}
