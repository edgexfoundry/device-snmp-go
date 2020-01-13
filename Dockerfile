#
# Copyright (C) 2018 IOTech Ltd
#
# SPDX-License-Identifier: Apache-2.0

ARG BASE=golang:1.11-alpine
FROM ${BASE} AS builder

ARG MAKE="make build"
ARG ALPINE_PKG_BASE="make git"
ARG ALPINE_PKG_EXTRA=""

# Replicate the APK repository override.
# If it is no longer necessary to avoid the CDN mirrors we should consider dropping this as it is brittle.
RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
# Install our build time packages.
RUN apk add --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

WORKDIR $GOPATH/src/github.com/edgexfoundry/device-snmp-go

COPY go.mod .
COPY Makefile .

RUN make update

COPY . .
RUN ${MAKE}

FROM alpine

# Make sure you change the cmd/res/configuration.toml port if you change it here
ENV APP_PORT=49993
EXPOSE $APP_PORT

ENV DEVICE_PORT=161
EXPOSE $DEVICE_PORT

COPY --from=builder /go/src/github.com/edgexfoundry/device-snmp-go/cmd /

ENTRYPOINT ["/device-snmp-go","--profile=docker","--confdir=/res","--registry=consul://edgex-core-consul:8500"]
