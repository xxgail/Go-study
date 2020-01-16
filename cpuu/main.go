package main

import (
	"cpuu.exe/pkg/sys"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	stats := &sys.Stats{}

	app := wails.CreateApp(&wails.AppConfig{
		Width:  400,
		Height: 400,
		Title:  "CPU",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(stats)
	app.Run()
}
