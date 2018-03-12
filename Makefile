.PHONY: all

all:
	go get -v -u github.com/labstack/echo
	go build -v -ldflags="-s -w"
