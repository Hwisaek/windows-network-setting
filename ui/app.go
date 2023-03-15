package ui

import (
	"os"

	"network-change/consts"
	"network-change/ui/layout"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var appInstance fyne.App

func InitApp() {
	os.Setenv("FYNE_FONT", `C:\Windows\Fonts\malgun.ttf`)
	appInstance = app.New()

	w := appInstance.NewWindow(consts.Title)
	w.Resize(fyne.NewSize(600, 700))

	grid := layout.GetBody(appInstance, w)
	w.SetContent(grid)

	w.ShowAndRun()
}
