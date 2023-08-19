.PHONY:build

build:
	go build -ldflags="-w -s"

install:
	mv -v dotsync /usr/bin

install_remote:
	go install github.com/mbaraa/dotsync@latest
