package layout

import (
	"errors"
	"network-change/status"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewFooter(w fyne.Window) *fyne.Container {
	apply := widget.NewButton("Apply", func() {
		ipAuto, err := status.IPAuto.Get()
		if err != nil {
			dialog.ShowInformation("error", "Failed to get IP mode", w)
		}

		interfaceName, err := status.NetworkAdapter.Get()
		if err != nil {
			dialog.ShowInformation("error", "Failed to get NetworkAdapter", w)
		}

		ipAddress := status.IP.Text
		subnetMask := status.SubnetMask.Text
		gateway := status.Gateway.Text
		dns := status.DNS.Text

		var cmd *exec.Cmd
		var cmd2 *exec.Cmd
		if ipAuto {
			cmd = exec.Command("netsh", "interface", "ip", "set", "address", interfaceName, "dhcp")
			cmd2 = exec.Command("netsh", "interface", "ip", "set", "dns", interfaceName, "dhcp")
		} else {
			cmd = exec.Command("netsh", "interface", "ip", "set", "address", interfaceName, "static", ipAddress, subnetMask, gateway)
			cmd2 = exec.Command("netsh", "interface", "ip", "set", "dns", interfaceName, "static", dns, "primary")
		}

		if err := cmd.Run(); err != nil {
			dialog.ShowError(errors.Join(err, errors.New("failed to change IP mode")), w)
			return
		}

		if err := cmd2.Run(); err != nil {
			dialog.ShowError(errors.Join(err, errors.New("failed to change DNS mode")), w)
			return
		}

		dialog.ShowInformation("Success", "IP mode changed successfully.", w)
	})
	return container.New(layout.NewGridLayout(1), apply)
}
