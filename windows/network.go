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
	DHCPEnabled          bool
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
	IPAutoMode           bool
}

func GetNetworkAdapters() ([]networkAdapter, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get network interfaces: %s", err)
	}

	adapters := make([]networkAdapter, 0, len(interfaces))
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, fmt.Errorf("failed to get network addresses: %s", err)
		}
		if len(addrs) == 0 {
			continue // no address found
		}
		adapter := networkAdapter{
			index: iface.Index,
			Name:  iface.Name,
			MAC:   strings.ToUpper(iface.HardwareAddr.String()),
		}
		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil || ip.IsLoopback() {
				continue
			}
			if ip.To4() != nil {
				adapter.IP = ip.String()
				break
			}
		}
		if adapter.IP == "" {
			continue // no IPv4 address found
		}
		adapters = append(adapters, adapter)
	}

	var configs []Win32_NetworkAdapterConfiguration
	err = wmi.Query("SELECT * FROM Win32_NetworkAdapterConfiguration", &configs)
	if err != nil {
		return nil, fmt.Errorf("failed to get network configuration: %s", err)
	}
	for i := range adapters {
		for _, config := range configs {
			if strings.ToUpper(config.MACAddress) == adapters[i].MAC {
				adapters[i].IPSubnet = config.IPSubnet
				adapters[i].DefaultIPGateway = config.DefaultIPGateway
				adapters[i].DNSServerSearchOrder = config.DNSServerSearchOrder
				adapters[i].IPAutoMode = config.DHCPEnabled
				break
			}
		}
	}
	return adapters, nil
}

func GetNetworks() ([]string, error) {
	networkAdapters, err := GetNetworkAdapters()
	if err != nil {
		return nil, err
	}
	networks := make([]string, len(networkAdapters))
	for i, adapter := range networkAdapters {
		networks[i] = adapter.Name
	}
	return networks, nil
}
