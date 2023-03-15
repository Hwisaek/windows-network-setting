package windows

import (
	"fmt"
	"net"
	"strings"

	"github.com/yusufpapurcu/wmi"
)

type Win32_NetworkAdapterConfiguration struct {
	Index                int
	IPSubnet             []string
	DefaultIPGateway     []string
	DNSServerSearchOrder []string
	MACAddress           string
	IPAddress            []string
}

type networkAdapter struct {
	index                int
	Name                 string
	IP                   string
	MAC                  string
	Description          string
	IPSubnet             []string
	DefaultIPGateway     []string
	DNSServerSearchOrder []string
}

func GetNetworkAdapters() (newNetworkAdapters []networkAdapter, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get network interfaces: %s", err)
	}

	for _, iface := range interfaces {
		i := networkAdapter{
			index: iface.Index,
			Name:  iface.Name,
			MAC:   strings.ToUpper(iface.HardwareAddr.String()),
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		if len(addrs) > 0 {
			var ip net.IP
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}

				if ip != nil && !ip.IsLoopback() && ip.To4() != nil {
					i.IP = ip.String()
					break
				}
			}
		}

		if i.IP == "" {
			continue
		}
		newNetworkAdapters = append(newNetworkAdapters, i)
	}

	// Query the Win32_NetworkAdapterConfiguration class to get network information
	var configs []Win32_NetworkAdapterConfiguration
	err = wmi.Query("SELECT * FROM Win32_NetworkAdapterConfiguration", &configs)
	if err != nil {
		return nil, fmt.Errorf("failed to get network configuration: %s", err)
	}

	// Print network information
	for _, config := range configs {
		for i := range newNetworkAdapters {
			if newNetworkAdapters[i].MAC == config.MACAddress {
				if len(config.IPSubnet) == 0 {
					newNetworkAdapters = append(newNetworkAdapters[:i], newNetworkAdapters[i+1:]...)
					break
				}
				newNetworkAdapters[i].IPSubnet = config.IPSubnet
				newNetworkAdapters[i].DefaultIPGateway = config.DefaultIPGateway
				newNetworkAdapters[i].DNSServerSearchOrder = config.DNSServerSearchOrder
				break
			}
		}
	}
	return
}
