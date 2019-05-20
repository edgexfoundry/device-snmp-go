.PHONY: build test clean prepare update docker

GO = CGO_ENABLED=0 GO111MODULE=on go

MICROSERVICES=cmd/device-snmp-switch-go

.PHONY: $(MICROSERVICES)

DOCKERS=docker_device_snmp_switch_go
.PHONY: $(DOCKERS)

VERSION=$(shell cat ./VERSION)
GIT_SHA=$(shell git rev-parse HEAD)
GOFLAGS=-ldflags "-X github.com/edgexfoundry-holding/device-snmp-switch-go.Version=$(VERSION)"

build: $(MICROSERVICES)
	$(GO) build ./...

cmd/device-snmp-switch-go:
	$(GO) build $(GOFLAGS) -o $@ ./cmd

test:
	$(GO) test ./... -cover

clean:
	rm -f $(MICROSERVICES)

run:
	cd bin && ./edgex-launch.sh

run_docker:
	cd bin && ./edgex-docker-launch.sh

docker: $(DOCKERS)

docker_device_snmp_switch_go:
	docker build \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry-holding/docker_device_snmp_switch_go:$(GIT_SHA) \
		-t edgexfoundry-holding/docker_device_snmp_switch_go:$(VERSION)-dev \
		.