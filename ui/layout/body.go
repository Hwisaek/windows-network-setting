package layout

import (
	"network-change/status"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewBody(appInstance fyne.App, w fyne.Window) *fyne.Container {
	header := NewHeader(appInstance, w)
	line1 := createLine("IP Address", status.IP)
	line2 := createLine("Subnetmask", status.SubnetMask)
	line3 := createLine("Gateway", status.Gateway)
	line4 := createLine("Default DNS", status.DNS)
	footer := NewFooter(w)

	return container.New(layout.NewGridLayout(1),
		header,
		line1, line2, line3, line4,
		footer,
	)
}

func createLine(labelText string, inputWidget fyne.CanvasObject) *fyne.Container {
	label := widget.NewLabel(labelText)
	form := &widget.Form{Items: []*widget.FormItem{{Widget: inputWidget}}}
	return container.New(layout.NewGridLayout(2), label, form)
}
