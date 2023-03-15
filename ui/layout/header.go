package layout

import (
	"network-change/status"
	"network-change/ui/component"
	"network-change/windows"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetHeader(appInstance fyne.App, w fyne.Window) *fyne.Container {
	networks, err := windows.GetNetworks()
	if err != nil {
		dialog.ShowInformation("error", "Failed to get network adapters.", w)
	}
	selectNetwork := widget.NewSelect(networks, func(s string) {
		adapters, _ := windows.GetNetworkAdapters()
		for _, v := range adapters {
			if v.Name == s {
				status.IP.SetText(v.IP)
				if len(v.IPSubnet) > 0 {
					status.Subnetmask.SetText(v.IPSubnet[0])
				}
				if len(v.DefaultIPGateway) > 0 {
					status.Gateway.SetText(v.DefaultIPGateway[0])
				}
				if len(v.DNSServerSearchOrder) > 0 {
					status.DNS.SetText(v.DNSServerSearchOrder[0])
				}
			}
		}
	})
	checkAuto := widget.NewCheck("Auto", func(b bool) {})
	btn1 := widget.NewButton("Network adapters", func() {
		component.ShowNetworkAdapters(appInstance, w)
		selectNetwork.Refresh()
	})
	return container.New(layout.NewAdaptiveGridLayout(3), selectNetwork, checkAuto, btn1)
}
