package network

import (
	"bytes"
	"net"
	"sort"
)

func InterfaceAddresses() ([]string, error) {
	iAddress, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	address := make([]string, len(iAddress))
	for i, addr := range iAddress {
		address[i] = addr.String()
	}
	return address, nil
}

func IsAddressInNetwork(address string, network string) bool {
	ip := net.ParseIP(address)
	_, subnet, err := net.ParseCIDR(network)
	if err != nil {
		return false
	}
	return subnet.Contains(ip)
}

func SortIPAddresses(ips []string) {
	sort.Slice(ips, func(i, j int) bool {
		return bytes.Compare(net.ParseIP(ips[i]), net.ParseIP(ips[j])) < 0
	})
}
