// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
	"github.com/edgexfoundry/device-snmp-go/internal/driver"
)

const (
	version     string = "1.0.0"
	serviceName string = "device-snmp-go"
)

func main() {
	sd := driver.SNMPDriver{}
	startup.Bootstrap(serviceName, version, &sd)
}
