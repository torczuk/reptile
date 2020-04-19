package network

import (
	"bytes"
	"fmt"
	"log"
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

func IsNetworkIpAddress(address string, network string) bool {
	ip := net.ParseIP(address)
	networkIp, _, err := net.ParseCIDR(network)
	if err != nil {
		return false
	}
	return ip.Equal(networkIp)
}

func SortIPAddresses(ips []string) {
	sort.Slice(ips, func(i, j int) bool {
		return bytes.Compare(net.ParseIP(ips[i]), net.ParseIP(ips[j])) < 0
	})
}

func MyAddress(ips []string) (int, error) {
	addresses, err := InterfaceAddresses()
	if err != nil {
		log.Fatalf("can't obtain interface addresses %v", err)
	}
	for i, ip := range ips {
		for _, add := range addresses {
			if IsNetworkIpAddress(ip, add) {
				return i, nil
			}
		}
	}
	return 0, fmt.Errorf("address not on the list %v", ips)
}
