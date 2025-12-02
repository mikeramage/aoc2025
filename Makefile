.DEFAULT_GOAL := build
PHONY: fmt vet build clean

clean: 
	rm aoc2025

fmt: 
	go fmt .

vet: fmt
	go vet .

build: vet
	go build .

