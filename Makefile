SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=java-folders-cleaner
PWD=$(shell pwd)

VERSION=1.0.0
BUILD_TIME=$(date "%FT%T%z")

LDFLAGS=-ldflags "-d -s -w -X tensin.org/clea,er/core/version.Build=`git rev-parse HEAD`" -a -tags netgo -installsuffix netgo
PACKAGE=tensin.org/cleaner

$(BINARY): 
	go build ${LDFLAGS} -o bin/${BINARY} ${PACKAGE}

.PHONY: install clean deploy run 

quick:
	go build -o bin/${BINARY} ${PACKAGE}

build:
	time go install ${PACKAGE}

install:
	time go install ${LDFLAGS} ${PACKAGE}

clean:
	[ -f bin/${BINARY} ] && rm -f bin/${BINARY}

run:
	bin/cleaner

init:
	# [ ! -f bin/glide ] && curl glide.sh/get | sh
	glide update
	glide install
	
test:
	go test -v tensin.org/cleaner/core/utils

docker:
	docker run --rm -it -v ${PWD}:/go tensin-app-golang /bin/bash

