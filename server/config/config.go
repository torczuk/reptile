package config

import (
	"os"
	"strings"
)

const Port int32 = 2600

func Servers() []string {
	serverEnv := os.Getenv("SERVERS")
	servers := strings.Split(serverEnv, " ")
	var result []string
	for i := range servers {
		trimmed := strings.Trim(servers[i], " ")
		if len(trimmed) != 0 {
			result = append(result, trimmed)
		}
	}
	return result
}
