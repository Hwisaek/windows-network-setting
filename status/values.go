package status

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var IP = widget.NewEntry()
var SubnetMask = widget.NewEntry()
var Gateway = widget.NewEntry()
var DNS = widget.NewEntry()

var IPAuto = binding.NewBool()
var NetworkAdapter = binding.NewString()

func init() {
	IPAuto.AddListener(binding.NewDataListener(func() {
		isAuto, _ := IPAuto.Get()
		if isAuto {
			IP.Disable()
			SubnetMask.Disable()
			Gateway.Disable()
			DNS.Disable()
		} else {
			IP.Enable()
			SubnetMask.Enable()
			Gateway.Enable()
			DNS.Enable()
		}
	}))
}
