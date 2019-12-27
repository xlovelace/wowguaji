package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	logEntry := widget.NewMultiLineEntry()
	logEntry.Disable()
	quit := make(chan bool)
	start := func() {
		text := logEntry.Text + "\r\n" + "start..."
		logEntry.SetText(text)
		go func() {
			for {
				select {
				case <-quit:
					return
				default:
					text = logEntry.Text + "\r\n" + "挂机中。。。"
					logEntry.SetText(text)
					time.Sleep(1 * time.Second)
				}
			}
		}()
	}
	stop := func() {
		text := logEntry.Text + "\r\n" + "stop..."
		logEntry.SetText(text)
		quit <- true
	}
	startButton := widget.NewButton("start", start)
	stopButton := widget.NewButton("stop", stop)
	w.SetContent(widget.NewVBox(
		widget.NewHBox(
			startButton,
			stopButton,
		),
		widget.NewScrollContainer(widget.NewVBox(logEntry)),
	),
	)
	w.ShowAndRun()

}

func start(e *widget.Entry) {
	fmt.Println("start...")
	e.SetText("start...")
}

func stop(e *widget.Entry) {
	fmt.Println("stop...")
	e.SetText("stop...")
}
