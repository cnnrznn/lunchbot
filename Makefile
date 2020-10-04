all: build

build: *.go
	git pull
	go build

run: build
	sudo setcap CAP_NET_BIND_SERVICE=+eip lunchbot
	./lunchbot