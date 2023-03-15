package ui

import (
	"os"

	"network-change/consts"
	"network-change/windows"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var App fyne.App

func getNetworksName(w fyne.Window) (networks []string) {
	networkAdapters, err := windows.GetNetworkAdapters()
	if err != nil {
		dialog.ShowInformation("error", "Failed to get network adapters.", w)
	}
	for _, v := range networkAdapters {
		networks = append(networks, v.Name)
	}
	return
}

func InitApp() {
	os.Setenv("FYNE_FONT", `C:\Windows\Fonts\malgun.ttf`)
	App := app.New()

	w := App.NewWindow(consts.Title)
	w.Resize(fyne.NewSize(600, 700))

	inputIp := widget.NewEntry()
	inputSubnetmask := widget.NewEntry()
	inputGateway := widget.NewEntry()
	inputDNS := widget.NewEntry()

	networks := getNetworksName(w)
	selectNetwork := widget.NewSelect(networks, func(s string) {
		adapters, _ := windows.GetNetworkAdapters()
		for _, v := range adapters {
			if v.Name == s {
				inputIp.SetText(v.IP)
				if len(v.IPSubnet) > 0 {
					inputSubnetmask.SetText(v.IPSubnet[0])
				}
				if len(v.DefaultIPGateway) > 0 {
					inputGateway.SetText(v.DefaultIPGateway[0])
				}
				if len(v.DNSServerSearchOrder) > 0 {
					inputDNS.SetText(v.DNSServerSearchOrder[0])
				}
			}
		}
	})

	checkAuto := widget.NewCheck("Auto", func(b bool) {
		switch b {
		case false:
			inputIp.Enable()
			inputSubnetmask.Enable()
			inputGateway.Enable()
			inputDNS.Enable()
		case true:
			inputIp.Disable()
			inputSubnetmask.Disable()
			inputGateway.Disable()
			inputDNS.Disable()
		}
	})
	btn1 := widget.NewButton("Network adapters", func() {
		ShowNetworkAdapters(w)
		selectNetwork.Refresh()
	})

	header := container.New(layout.NewAdaptiveGridLayout(3), selectNetwork, checkAuto, btn1)

	form1 := &widget.Form{Items: []*widget.FormItem{{Widget: inputIp}}}

	line1 := container.New(layout.NewGridLayout(2), widget.NewLabel("IP Address"), form1)

	form2 := &widget.Form{Items: []*widget.FormItem{{Widget: inputSubnetmask}}}
	line2 := container.New(layout.NewGridLayout(2), widget.NewLabel("Subnetmask"), form2)

	form3 := &widget.Form{Items: []*widget.FormItem{{Widget: inputGateway}}}
	line3 := container.New(layout.NewGridLayout(2), widget.NewLabel("Gateway"), form3)

	form4 := &widget.Form{Items: []*widget.FormItem{{Widget: inputDNS}}}
	line4 := container.New(layout.NewGridLayout(2), widget.NewLabel("Default DNS"), form4)

	apply := widget.NewButton("Apply", func() {})
	buttons := container.New(layout.NewGridLayout(1), apply)

	grid := container.New(layout.NewGridLayout(1),
		header,
		line1, line2, line3, line4,
		buttons)

	w.SetContent(grid)

	w.ShowAndRun()
}
