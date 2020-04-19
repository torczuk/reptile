package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfig_ServerFromEnv(t *testing.T) {
	os.Setenv("SERVERS", "172.28.1.1 172.28.1.2  172.28.1.3 ")

	servers := Servers()

	assert.Equal(t, []string{"172.28.1.1", "172.28.1.2", "172.28.1.3"}, servers)
}
