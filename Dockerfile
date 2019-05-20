#
# Copyright (C) 2018 IOTech Ltd
#
# SPDX-License-Identifier: Apache-2.0

ARG ALPINE=golang:1.11-alpine
FROM ${ALPINE} AS builder
ARG ALPINE_PKG_BASE="build-base git openssh-client"
ARG ALPINE_PKG_EXTRA=""

# Replicate the APK repository override.
# If it is no longer necessary to avoid the CDN mirros we should consider dropping this as it is brittle.
RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories
# Install our build time packages.
RUN apk add --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

WORKDIR $GOPATH/src/github.com/edgexfoundry-holding/device-snmp-switch-go

COPY . .

# To run tests in the build container:
#   docker build --build-arg 'MAKE=build test' .
# This is handy of you do your Docker business on a Mac
ARG MAKE=build
RUN make $MAKE


FROM scratch

# Make sure you change the cmd/res/configuration.toml port if you change it here
ENV APP_PORT=49993
EXPOSE $APP_PORT

ENV DEVICE_PORT=161
EXPOSE $DEVICE_PORT

COPY --from=builder /go/src/github.com/edgexfoundry-holding/device-snmp-switch-go/cmd /

ENTRYPOINT ["/device-snmp-switch-go","--registry","--profile=docker","--confdir=/res"]