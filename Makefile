.PHONY : all fmt test

all : test
	CGO_ENABLED=0 go install -a -v

test : fmt
	go test -v ./...
	cd e2e && bats *.bats

fmt :
	go fmt ./...
	