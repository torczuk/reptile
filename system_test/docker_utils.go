package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
)

func main() {

}

func GetContainerIdByName(containerName string) ([]string, error) {
	containers, err := ListContainers()
	if err != nil {
		return make([]string, 0), err
	}
	ids := make([]string, 0)
	for _, container := range containers {
		for _, name := range container.Names {
			if strings.Compare(containerName, strings.TrimLeft(name, "/")) == 0 {
				ids = append(ids, container.ID)
			}
		}
	}
	return ids, nil
}

func ListContainers() ([]types.Container, error) {
	docker, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return docker.ContainerList(context.Background(), types.ContainerListOptions{All: false})
}