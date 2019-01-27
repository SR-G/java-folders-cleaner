SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
SHELL=bash

BINARY=cleaner
PWD=$(shell pwd)

VERSION=1.0.0
BUILD_TIME=$(date "%FT%T%z")

LDFLAGS=-ldflags "-d -s -w -X tensin.org/cleaner/core/version.Build=`git rev-parse HEAD`" -a -tags netgo -installsuffix netgo
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
	-@rm -f bin/*.zip 2>/dev/null || true

distribution: clean install
	mkdir /go/bin/linux/ 
	mv /go/bin/${BINARY} /go/bin/linux/
	cp /go/resources/cleaner.conf /go/bin/linux/
	cp /go/resources/cleaner.conf /go/bin/windows_amd64/
	cd /go/bin/ ; zip -r -9 ${BINARY}.zip ./linux ; zip -r -9 ${BINARY}.zip ./windows_amd64

test:
	-@mkdir -p bin/eclipse_installations/{photon,luna,mars,2018-12}/configuration/org.eclipse.osgi/{01,23,39,49,58,104} 2>/dev/null || true
	-@mkdir -p bin/workspaces/{workspace_a,workspace_b}/.metadata 2>/dev/null || true
	-@mkdir -p bin/projects/{project_a,project_b,project_c}/{bin,build,dist,target} 2>/dev/null || true
	-@touch bin/projects/{project_a,project_b}/target/{dependency1,dependency2,dependency3}.jar || true
	-@touch bin/projects/{project_a,project_b,project_c}/.{project,classpath} || true
	-@touch bin/projects/{project_a,project_b,project_c}/bin/{file1,file2,file3,file4,file5}.class || true
	-@touch bin/eclipse_installations/{photon,luna,mars,2018-12}/eclipse.exe 2>/dev/null || true

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

