.PHONY : all test e2e

all : e2e
	@echo Everything is OK

e2e :
	cd e2e && bats *.bats

install : test
	CGO_ENABLED=0 go install -a -v

test :
	go test -v ./...

fmt :
	go fmt ./...
	
