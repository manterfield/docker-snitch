package main

import (
    "encoding/json"
    "fmt"
    "github.com/manterfield/docker-snitch"
    docker "github.com/fsouza/go-dockerclient"
    "net/http"
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
        json.NewEncoder(w).Encode(rundeckContainerOpts(containers))
    })
}

type RundeckContainerOpt struct {
    Name string `json:"name"`
    Value string `json:"value"`
} 

func rundeckContainerOpts(ctrs []docker.APIContainers) []RundeckContainerOpt {
    var containerOpts []RundeckContainerOpt
    var ctrOpt RundeckContainerOpt
    var optName string

    for _, ctr := range ctrs {
        if ctr.Image == "tictocstech/docker-snitch" {
            continue    
        }
        optName = strings.Join(append(ctr.Names, ctr.Image), " ")
        optName = strings.TrimLeft(optName, "/")
        ctrOpt = RundeckContainerOpt{Name: optName, Value: ctr.ID}
        containerOpts = append(containerOpts, ctrOpt)
    }

    return containerOpts
}

func main() {
    http.Handle("/", index())
    http.Handle("/containers", containers())
    http.Handle("/container-opts", containerOpts())
    http.ListenAndServe(":8080", nil)
}
