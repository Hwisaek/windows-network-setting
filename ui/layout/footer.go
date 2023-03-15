package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetFooter() *fyne.Container {
	apply := widget.NewButton("Apply", func() {})
	return container.New(layout.NewGridLayout(1), apply)
}
