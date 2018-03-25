SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=cleaner
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
	GOARCH=amd64 GOOS=linux go install ${LDFLAGS} ${PACKAGE}
	GOARCH=amd64 GOOS=windows go install ${PACKAGE}

clean:
	-@rm -f bin/${BINARY} 2>/dev/null || true
	-@rm -rf bin/linux 2>/dev/null || true
	-@rm -rf bin/windows_amd64 2>/dev/null || true

distribution: clean install
	mkdir /go/bin/linux/ 
	mv /go/bin/${BINARY} /go/bin/linux/
	# cp /go/src/main/resources/sol.json /go/bin/linux/ 
	# cp /go/src/main/resources/sol.json /go/bin/windows_amd64/
	# cp /go/src/script/*.bat /go/bin/windows_amd64
	cd /go/bin/ ; zip -r -9 ${BINARY}.zip ./linux ; zip -r -9 ${BINARY}.zip ./windows_amd64

test:
	-@mkdir bin/target 2>/dev/null || true
	-@mkdir bin/bin 2>/dev/null || true
	-@touch bin/bin/test.class 2>/dev/null || true
	-@touch bin/target/test.jar 2>/dev/null || true
	-@touch bin/target/test.log 2>/dev/null || true
	-@touch bin/test.log 2>/dev/null || true
	-@touch bin/toto.class 2>/dev/null || true
	-@mkdir bin/logs 2>/dev/null || true

run:
	bin/cleaner

init:
	# [ ! -f bin/glide ] && curl glide.sh/get | sh
	glide update
	glide install
	
# test:
# 	go test -v tensin.org/cleaner/core/utils

docker:
	docker run --rm -it -v ${PWD}:/go tensin-app-golang /bin/bash

