package main

import (
	"go-fileserver/file"
)

func main() {
	file.ServeHTTP()
	file.StartFileServer()
	//a := app.New()
	//w := a.NewWindow("Hello")
	//
	//hello := widget.NewLabel("Hello Fyne!")
	//w.SetContent(widget.NewVBox(
	//	hello,
	//	widget.NewButton("Hi!", func() {
	//		hello.SetText("Welcome :)")
	//	}),
	//))
	//
	//w.ShowAndRun()
}