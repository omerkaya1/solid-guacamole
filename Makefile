.DEFAULT_GOAL := build

build:
	go build -ldflags "-s -w -extldflags" -o solid-gaucamole .

kill:
	kill -s SIGKILL $(pgrep solid-guacamole)

ack:
	kill -s SIGUSR1 $(pgrep solid-guacamole)