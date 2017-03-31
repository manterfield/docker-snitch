# docker-snitch
Docker Snitch **BETA**

A service exposing info on running containers over a simple web API. Can be used as a remote option provider to Rundeck.

Currently exposes two endpoints:
`/containers` and `/container-opts`

Easiest is to run as a docker container, like so:
`$ docker run -d -p 8080:8080 -v /var/run/docker.sock:/tmp/docker.sock:ro tictocstech/docker-snitch`

### TODO
- [x] Create Dockerfile
- [x] Create simple webservice
- [x] Get info on running containers
- [x] Present info on running containers via webservice
- [ ] Add token auth
- [ ] Add monitoring metrics
