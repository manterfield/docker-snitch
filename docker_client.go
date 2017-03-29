package dockersnitch

import (
	docker "github.com/fsouza/go-dockerclient"
	"os"
)

func Containers(status string) []docker.APIContainers {
	endpoint := os.Getenv("DOCKER_ENDPOINT")
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	var containerOpts = docker.ListContainersOptions{All: true}

	switch status {
    case "created", "restarting", "running", "paused", "exited", "dead":
    	containerOpts.Filters = map[string][]string {
    		"status": []string{status},
    	}
    }

	ctrs, err := client.ListContainers(containerOpts)
	if err != nil {
		panic(err)
	}

	return ctrs
}
