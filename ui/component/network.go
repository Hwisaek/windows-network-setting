package component

import (
	"network-change/windows"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowNetworkAdapters(appInstance fyne.App, parent fyne.Window) {
	NetworkAdapters, err := windows.GetNetworkAdapters()
	if err != nil {
		dialog.ShowInformation("error", "Failed to get network adapters.", appInstance.NewWindow("error"))
		return
	}

	list := widget.NewTable(
		func() (int, int) {
			return len(NetworkAdapters) + 1, 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Row == 0 {
				switch i.Col {
				case 0:
					o.(*widget.Label).SetText("Name")
				case 1:
					o.(*widget.Label).SetText("IP Address")
				case 2:
					o.(*widget.Label).SetText("MAC Address")
				}
				return
			}

			switch i.Col {
			case 0:
				o.(*widget.Label).SetText(NetworkAdapters[i.Row-1].Name)
			case 1:
				o.(*widget.Label).SetText(NetworkAdapters[i.Row-1].IP)
			case 2:
				o.(*widget.Label).SetText(NetworkAdapters[i.Row-1].MAC)
			}
		})
	list.SetColumnWidth(0, 150)
	list.SetColumnWidth(1, 150)

	window := dialog.NewCustom("NetworkAdapters", "test", list, parent)
	window.Resize(fyne.NewSize(500, 400))
	window.Show()
}
