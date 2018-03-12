.PHONY: all

all:
	go build -v -ldflags="-s -w"
