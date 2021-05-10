// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 Dell Technologies
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"errors"
	"fmt"

	g "github.com/soniah/gosnmp"
)

// SNMPClient represents the SNMP device is used for getting and setting SNMP device via OID
type SNMPClient struct {
	ipAddr string
	ipPort uint16
}

func NewSNMPClient(addr string, port uint16) *SNMPClient {
	return &SNMPClient{
		ipAddr: addr,
		ipPort: port,
	}
}

type DeviceCommand struct {
	operation string
	value     int
}

func NewGetDeviceCommand(op string) DeviceCommand {
	return DeviceCommand{
		operation: op,
		// for Gets, value is not used
		value: 0,
	}
}

func NewSetDeviceCommand(op string, val int) DeviceCommand {
	return DeviceCommand{
		operation: op,
		value:     val,
	}
}

func (c *SNMPClient) GetValues(commands []DeviceCommand) ([]interface{}, error) {

	var results []interface{}
	var oids []string

	for _, command := range commands {
		if command.operation == "" {
			return results, errors.New("Unknown operation: " + command.operation)
		}
		oids = append(oids, command.operation)
	}

	if g.Default.Conn == nil {
		g.Default.Target = c.ipAddr
		g.Default.Port = c.ipPort
		g.Default.Community = COMMUNITY_ACCESS
		err := g.Default.Connect()
		if err != nil {
			return results, err
		}
	}
	packets, err2 := g.Default.Get(oids)

	if err2 != nil {
		return results, err2
	}

	for _, variable := range packets.Variables {
		switch variable.Type {
		case g.OctetString:
			var buf = variable.Value.([]byte)
			var octet = string(buf)
			if len(buf) == 6 {
				octet = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
			}
			results = append(results, string(octet))
		case g.IPAddress:
			results = append(results, string(variable.Value.(string)))
		default:
			var temp = g.ToBigInt(variable.Value).Uint64()
			results = append(results, int(temp))
		}
	}
	return results, nil
}

func (c *SNMPClient) GetValue(command DeviceCommand) (interface{}, error) {
	commands := []DeviceCommand{command}
	results, err := c.GetValues(commands)
	if err != nil {
		return 0, err
	}
	return results[0], nil
}

func (c *SNMPClient) SetValues(commands []DeviceCommand) ([]interface{}, error) {

	var results []interface{}
	var pdus []g.SnmpPDU

	if g.Default.Conn == nil {
		g.Default.Target = c.ipAddr
		g.Default.Port = c.ipPort
		g.Default.Community = COMMUNITY_ACCESS
		err := g.Default.Connect()
		if err != nil {
			return results, err
		}
	}

	for _, command := range commands {
		if command.operation == "" {
			return results, errors.New("Unknown operation: " + command.operation)
		}
		// TODO pass in logger
		pdu := g.SnmpPDU{Name: command.operation, Type: g.Integer, Value: command.value, Logger: nil}
		pdus = append(pdus, pdu)
	}

	packets, err2 := g.Default.Set(pdus)
	if err2 != nil {
		return results, err2
	}

	for _, variable := range packets.Variables {
		switch variable.Type {
		case g.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
			var octet = string(variable.Value.([]byte))
			results = append(results, string(octet))
		default:
			var temp = g.ToBigInt(variable.Value).Uint64()
			results = append(results, int(temp))
		}
	}
	return results, nil
}

func (c *SNMPClient) Disconnect() {
	if g.Default.Conn != nil {
		g.Default.Conn.Close()
	}
}
