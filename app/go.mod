module github.com/vextor22/go_docker

go 1.12

replace github.com/vextor22/go_docker/app/restservice => ./restservice

require (
	github.com/gorilla/mux v1.7.1
	github.com/vextor22/go_docker/app/restservice v0.0.0-00010101000000-000000000000
)
