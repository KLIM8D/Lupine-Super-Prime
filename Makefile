# Author: https://github.com/jpoehls 
# Modified by: klim8d

.PHONY: build doc fmt lint run debug test vendor_clean vendor_get vendor_update vet

APP  = Lupine-Super-Prime
DEPS_FOLDER = .vendor
DEPS = github.com/garyburd/redigo/redis

# Prepend our _vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
GP := ${PWD}/$(DEPS_FOLDER):${GOPATH}
export GOPATH=$(GP)

default: build

build: vet
	go build -v -o ./bin/$(APP) ./

doc:
	godoc -http=:6060 -index

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./...

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint ./

run:
	./bin/$(APP)

debug:
	go run ./main.go -mode=0

test:
	go test -v ./...

benchmark:
	go test -v -run=XXX -bench=. ./...

vendor_clean:
	rm -dRf ./$(DEPS_FOLDER)/

# We have to set GOPATH to just the .vendor
# directory to ensure that `go get` doesn't
# update packages in our primary GOPATH instead.
# This will happen if you already have the package
# installed in GOPATH since `go get` will use
# that existing location as the destination.
vendor_get: vendor_clean
	GOPATH=${PWD}/$(DEPS_FOLDER) go get -d -u -v \
    $(DEPS)

vendor_update: vendor_get
	rm -rf `find ./$(DEPS_FOLDER)/src -type d -name .git` \
    && rm -rf `find ./$(DEPS_FOLDER)/src -type d -name .hg` \
    && rm -rf `find ./$(DEPS_FOLDER)/src -type d -name .bzr` \
    && rm -rf `find ./$(DEPS_FOLDER)/src -type d -name .svn`

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
vet:
	go vet ./...
