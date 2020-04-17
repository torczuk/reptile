package network

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAddressInNetwork_InSubnet(t *testing.T) {
	address := "192.168.1.104"
	subnet := "192.168.1.104/24"

	inSubnet := IsAddressInNetwork(address, subnet)
	assert.True(t, inSubnet, fmt.Sprintf("address %v should be in %v", address, subnet))
}

func TestIsAddressInNetwork_OutsideSubnet(t *testing.T) {
	address := "192.168.0.104"
	subnet := "192.168.1.104/24"

	inSubnet := IsAddressInNetwork(address, subnet)
	assert.False(t, inSubnet, fmt.Sprintf("address %v shouldn't be in %v", address, subnet))
}

func TestInterfaceAddresses(t *testing.T) {
	addresses, err := InterfaceAddresses()

	assert.Nil(t, err)
	assert.NotEmpty(t, addresses)
}

func TestSortIPAddresses(t *testing.T) {
	ips := []string{"192.168.1.2", "192.168.1.1", "192.168.2.1", "10.152.16.23", "69.52.220.44"}

	SortIPAddresses(ips)

	assert.Equal(t, ips, []string{"10.152.16.23", "69.52.220.44", "192.168.1.1", "192.168.1.2", "192.168.2.1"})
}
