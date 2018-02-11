#Tracer TCP server

Further Tracer API server

##Install
 - `git clone https://github.com/morrah77/tracer.git`

##Ensure dependencies
 - `./prepare.sh` (please don't forget to check your GOPATH to fit go dep[https://golang.github.io/dep/] requirements)
 
##Build
 - Locally: `./build.sh`
 - Docker image: `./build.sh docker`
 
##Run
 - Locally: `./bin/server -listen 8088`
 - In Docker container: `./run.sh`

#See server logs
 - `./log.sh`

##Test
 - Manually: `curl -iv <Container IP address>:8088/my/path/1 -d 12345abcdef`
 - Auto: `./test.sh [package[{ package}]]`

##Stop docker container:
 - `./stop.sh`
