// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"github.com/edgexfoundry-holding/device-snmp-switch-go/internal/driver"
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
)

const (
	version     string = "1.0.0"
	serviceName string = "device-snmp-switch-go"
)

func main() {
	sd := driver.SNMPDriver{}
	startup.Bootstrap(serviceName, version, &sd)
}
