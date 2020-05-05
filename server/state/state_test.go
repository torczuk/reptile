package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyIp_WhenDefined(t *testing.T) {
	state := &ReplicaState{Configuration: []string{"192.168.1.2", "192.168.1.1"}, MyAddress: 1}

	myIp := state.MyIp()

	assert.Equal(t, "192.168.1.1", myIp)
}

func TestMyIp_WhenNotDefined(t *testing.T) {
	state := &ReplicaState{Configuration: []string{}, MyAddress: 0}

	myIp := state.MyIp()

	assert.Equal(t, "", myIp)
}

func TestOthersIp_WhenNotDefined(t *testing.T) {
	state := &ReplicaState{Configuration: []string{}, MyAddress: 0}

	others := state.OthersIp()

	assert.Empty(t, others)
}

func TestOthersIp_WhenDefined(t *testing.T) {
	state := &ReplicaState{Configuration: []string{"192.168.1.2", "192.168.1.1", "192.168.1.0"}, MyAddress: 1}

	others := state.OthersIp()

	assert.Equal(t, []string{"192.168.1.2", "192.168.1.0"}, others)
}

func TestAmIPrimary_WhenPrimary(t *testing.T) {
	state := &ReplicaState{Configuration: []string{"192.168.1.2", "192.168.1.1", "192.168.1.0"}, MyAddress: 1}

	primary := state.AmIPrimary()

	assert.Equal(t, false, primary, "replica can't be primary")
}

func TestAmIPrimary_WhenBackup(t *testing.T) {
	state := &ReplicaState{Configuration: []string{"192.168.1.2", "192.168.1.1", "192.168.1.0"}, MyAddress: 0}

	primary := state.AmIPrimary()

	assert.Equal(t, true, primary, "replica must be primary")
}

func TestAmIPrimary_WhenNotDefined(t *testing.T) {
	state := &ReplicaState{Configuration: []string{}, MyAddress: 0}

	primary := state.AmIPrimary()

	assert.Equal(t, false, primary, "unknown primary/backup")
}