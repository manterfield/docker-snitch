> I haven't touched this project or used it in years. I'd recommend not using it directly, but if you want an example of exposing data to Rundeck or hooking into the docker daemon API, then it should provide a simple enough example.

# docker-snitch
Docker Snitch **BETA**

A service exposing info on running containers over a simple web API. Can be used as a remote option provider to Rundeck.

Currently exposes two endpoints:
`/containers` and `/container-opts`

Easiest is to run as a docker container, like so:
`$ docker run -d -p 8080:8080 -v /var/run/docker.sock:/tmp/docker.sock:ro tictocstech/docker-snitch`
