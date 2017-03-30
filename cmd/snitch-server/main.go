package main

import (
    "encoding/json"
    "fmt"
    "github.com/manterfield/docker-snitch"
    docker "github.com/fsouza/go-dockerclient"
    "net/http"
    "sort"
    "strings"
)

func index() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world!")
    })
}

func containers() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        status := r.URL.Query().Get("status")
        containers := dockersnitch.Containers(status)

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(containers)
    })
}

func containerOpts() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        status := r.URL.Query().Get("status")
        containers := dockersnitch.Containers(status)

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(getContainerOpts(containers))
    })
}

type ContainerOpt struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

func getContainerOpts(ctrs []docker.APIContainers) []ContainerOpt {
    var containerOpts []ContainerOpt
    var ctrOpt ContainerOpt
    var optName string

    for _, ctr := range ctrs {
        if ctr.Image == "tictocstech/docker-snitch" {
            continue    
        }
        optName = strings.Join(append(ctr.Names, ctr.Image), " ")
        optName = strings.TrimLeft(optName, "/")
        ctrOpt = ContainerOpt{Name: optName, Value: ctr.ID}
        containerOpts = append(containerOpts, ctrOpt)
    }

    sort.Slice(containerOpts, func(i, j int) bool {
        return containerOpts[i].Name < containerOpts[j].Name
    })
    return containerOpts
}

func main() {
    http.Handle("/", index())
    http.Handle("/containers", containers())
    http.Handle("/container-opts", containerOpts())
    http.ListenAndServe(":8080", nil)
}
