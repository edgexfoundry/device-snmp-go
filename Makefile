.PHONY: build test clean prepare update docker

GO=CGO_ENABLED=0 GO111MODULE=on go
GOCGO=CGO_ENABLED=1 GO111MODULE=on go

MICROSERVICES=cmd/device-snmp

.PHONY: $(MICROSERVICES)

DOCKERS=docker_device_snmp_go
.PHONY: $(DOCKERS)

# This pulls the version of the SDK from the go.mod file. It works by looking for the line
# with the SDK and printing just the version number that comes after it.
SDKVERSION=$(shell sed -En 's|.*github.com/edgexfoundry/device-sdk-go/v2 (v[\.0-9a-zA-Z-]+).*|\1|p' go.mod)

# this pulls the version from local VERSION file that is created by the Jenkins Pipeline.
VERSION=$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)

GIT_SHA=$(shell git rev-parse HEAD)
GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-snmp-go.Version=$(VERSION) \
                  -X github.com/edgexfoundry/device-sdk-go/v2/internal/common.SDKVersion=$(SDKVERSION)"

build: $(MICROSERVICES)

tidy:
	go mod tidy

cmd/device-snmp:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd

test:
	$(GOCGO) test ./... -coverprofile=coverage.out
	$(GOCGO) vet ./...
	gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")
	[ "`gofmt -l $$(find . -type f -name '*.go'| grep -v "/vendor/")`" = "" ]
	./bin/test-attribution-txt.sh

clean:
	rm -f $(MICROSERVICES)

update:
	$(GO) mod download

run:
	cd bin && ./edgex-launch.sh

docker: $(DOCKERS)

docker_device_snmp_go:
	docker build \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/device-snmp:$(GIT_SHA) \
		-t edgexfoundry/device-snmp:$(VERSION)-dev \
		.

vendor:
	$(GO) mod vendor
