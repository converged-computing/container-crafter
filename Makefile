HERE ?= $(shell pwd)
LOCALBIN ?= $(shell pwd)/bin

.PHONY: all

all: build

.PHONY: $(LOCALBIN)
$(LOCALBIN):
	mkdir -p $(LOCALBIN)
	
build: $(LOCALBIN)
	GO111MODULE="on" go build -o $(LOCALBIN)/container-crafter cmd/container-crafter/container-crafter.go

build-arm: $(LOCALBIN)
	GO111MODULE="on" GOARCH=arm64 go build -o $(LOCALBIN)/container-crafter-arm cmd/container-crafter/container-crafter.go

build-ppc: $(LOCALBIN)
	GO111MODULE="on" GOARCH=ppc64le go build -o $(LOCALBIN)/container-crafter-ppc cmd/container-crafter/container-crafter.go