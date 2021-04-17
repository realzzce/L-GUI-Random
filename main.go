package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
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
	var logstr string
	// default 0 - 99
	var Max, Min int
	Min = 0
	Max = 100
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Hi")
	w.Resize(fyne.NewSize(300, 150))

	inputMin := widget.NewEntry()
	inputMin.SetPlaceHolder("MIN")
	inputMax := widget.NewEntry()
	inputMax.SetPlaceHolder("MAX - 1")

	inputMin.OnChanged = func(content string) {
		fmt.Println(inputMin.Text)
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
	lognum := widget.NewLabel("")

	w.SetContent(widget.NewVBox(
		widget.NewHBox(
			inputMin,
			inputMax,
		),
		widget.NewSeparator(),
		hello,
		widget.NewSeparator(),
		widget.NewHBox(
			widget.NewButton("Random", func() {
				fmt.Println(Min, Max)
				randomtext := fmt.Sprint(random(Min, Max))
				hello.SetText(randomtext)
			}),
			widget.NewButton("Random * 10", func() {
				logstr = ""
				for i := 1; i <= 10; i++ {
					randomtext := fmt.Sprint(random(Min, Max))
					hello.SetText(randomtext)
					if i == 10 {
						logstr = logstr + randomtext
					} else {
						logstr = logstr + randomtext + " , "
					}
					lognum.SetText(logstr)
				}
			}),
			widget.NewButton("Clear", func() {
				hello.SetText("")
				lognum.SetText("")
			}),
		),
		lognum,
	),
	)

	w.ShowAndRun()
}
