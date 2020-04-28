// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 Dell Technologies
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"strconv"
	"sync"
	"time"
)

type SNMPDriver struct {
	lc           logger.LoggingClient
	asyncCh      chan<- *dsModels.AsyncValues
	switchButton bool
}

var client *SNMPClient

//Used to avoid get/set at the same time. If this happens simultaneously, state
//of the device can get out of sync with command actuation result
var mu sync.Mutex

// DisconnectDevice handles protocol-specific cleanup when a device
// is removed.
func (s *SNMPDriver) DisconnectDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	s.lc.Warn("Driver's DisconnectDevice function didn't implement")
	return nil
}

// Initialize performs protocol-specific initialization for the device
// service.
func (s *SNMPDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues, deviceCh chan<- []dsModels.DiscoveredDevice) error {
	s.lc = lc
	s.asyncCh = asyncCh
	return nil
}

// HandleReadCommands triggers a protocol Read operation for the specified device.
func (s *SNMPDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {

	res = make([]*dsModels.CommandValue, 0)

	if len(reqs) != 1 {
		err = fmt.Errorf("SimpleDriver.HandleReadCommands; too many command requests; only one supported")
		return
	}

	var TCP = protocols["TCP"]

	Name := deviceName
	Address := TCP["Address"]
	Port := TCP["Port"]

	//s.lc.Debug(fmt.Sprintf("SimpleDriver.HandleReadCommands: protocols: %v operation: %v attributes: %v", protocols, reqs[0].RO.Operation, reqs[0].DeviceResource.Attributes))

	s.lc.Debug("Port %s", Port)
	s.lc.Debug("Address", Address)

	var commands []DeviceCommand
	for _, req := range reqs {
		oid := req.Attributes["oid"]
		s.lc.Debug(fmt.Sprintf("SNMPDriver.HandleReadCommand: device: %s operation: %v attributes: %v", Name, req.DeviceResourceName, req.Attributes))

		commands = append(commands, NewGetDeviceCommand(oid))
	}

	port, err := strconv.ParseUint(Port, 10, 64)

	mu.Lock()
	if client == nil {
		client = NewSNMPClient(Address, uint16(port))
	}

	now := time.Now().UnixNano()

	// this an be any type either string or int at this point
	var vals, err2 = client.GetValues(commands)

	mu.Unlock()

	if err2 != nil {
		s.lc.Error(fmt.Sprintf("SNMPDriver.HandleReadCommands; %s", err2))
		return
	}

	for i, val := range vals {
		// switch here on type
		//s.lc.Debug(fmt.Sprintf("SNMPDriver.HandleReadCommands; value of %v is: %d", reqs[i].RO, val))

		switch val.(type) {
		case string:
			cv, _ := dsModels.NewCommandValue(reqs[i].DeviceResourceName, now, val, dsModels.String)
			res = append(res, cv)
		case int64:
			safeVal, ok := val.(int64) // Alt. non panicking version
			if ok {
				cv, _ := dsModels.NewInt64Value(reqs[i].DeviceResourceName, now, int64(safeVal))
				res = append(res, cv)
			} else {
				s.lc.Error(fmt.Sprintf("SNMPDriver.HandleWriteCommands; %s", err2))
			}
		case int32:
			safeVal, ok := val.(int32) // Alt. non panicking version
			if ok {
				cv, _ := dsModels.NewInt32Value(reqs[i].DeviceResourceName, now, int32(safeVal))
				res = append(res, cv)
			} else {
				s.lc.Error(fmt.Sprintf("SNMPDriver.HandleWriteCommands; %s", err2))
			}
		case int16:
			safeVal, ok := val.(int16) // Alt. non panicking version
			if ok {
				cv, _ := dsModels.NewInt16Value(reqs[i].DeviceResourceName, now, int16(safeVal))
				res = append(res, cv)
			} else {
				s.lc.Error(fmt.Sprintf("SNMPDriver.HandleWriteCommands; %s", err2))
			}
		case int8:
			safeVal, ok := val.(int8) // Alt. non panicking version
			if ok {
				cv, _ := dsModels.NewInt8Value(reqs[i].DeviceResourceName, now, int8(safeVal))
				res = append(res, cv)
			} else {
				s.lc.Error(fmt.Sprintf("SNMPDriver.HandleWriteCommands; %s", err2))
			}
		case uint64:
			safeVal, ok := val.(uint64) // Alt. non panicking version
			if ok {
				cv, _ := dsModels.NewUint64Value(reqs[i].DeviceResourceName, now, uint64(safeVal))
				res = append(res, cv)
			} else {
				s.lc.Error(fmt.Sprintf("SNMPDriver.HandleWriteCommands; %s", err2))
			}
		case uint32:
			safeVal, ok := val.(uint32) // Alt. non panicking version
			if ok {
				cv, _ := dsModels.NewUint32Value(reqs[i].DeviceResourceName, now, uint32(safeVal))
				res = append(res, cv)
			} else {
				s.lc.Error(fmt.Sprintf("SNMPDriver.HandleWriteCommands; %s", err2))
			}
		case uint16:
			safeVal, ok := val.(uint16) // Alt. non panicking version
			if ok {
				cv, _ := dsModels.NewUint16Value(reqs[i].DeviceResourceName, now, uint16(safeVal))
				res = append(res, cv)
			} else {
				s.lc.Error(fmt.Sprintf("SNMPDriver.HandleWriteCommands; %s", err2))
			}
		case uint8:
			safeVal, ok := val.(uint8) // Alt. non panicking version
			if ok {
				cv, _ := dsModels.NewUint8Value(reqs[i].DeviceResourceName, now, uint8(safeVal))
				res = append(res, cv)
			} else {
				s.lc.Error(fmt.Sprintf("SNMPDriver.HandleWriteCommands; %s", err2))
			}
		case uint:
			safeVal, ok := val.(uint64) // Alt. non panicking version
			if ok {
				cv, _ := dsModels.NewUint64Value(reqs[i].DeviceResourceName, now, uint64(safeVal))
				res = append(res, cv)
			} else {
				s.lc.Error(fmt.Sprintf("SNMPDriver.HandleWriteCommands; %s", err2))
			}
		default:
			safeVal, ok := val.(int) // Alt. non panicking version
			if ok {
				cv, _ := dsModels.NewInt32Value(reqs[i].DeviceResourceName, now, int32(safeVal))
				res = append(res, cv)
			} else {
				s.lc.Error(fmt.Sprintf("SNMPDriver.HandleWriteCommands; %s", err2))
			}
		}
	}

	return
}

// HandleWriteCommands passes a slice of CommandRequest struct each representing
// a ResourceOperation for a specific device resource (aka DeviceObject).
// Since the commands are actuation commands, params provide parameters for the individual
// command.
func (s *SNMPDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest, params []*dsModels.CommandValue) error {

	var Address = "192.168.150.2"
	var Port = 161
	//var Name = "dellswitch01"

	var commands []DeviceCommand
	for i, req := range reqs {
		//s.lc.Debug(fmt.Sprintf("SNMPDriver.HandleWriteCommands: device: %s operation: %v attributes: %v", Name, req.RO.Operation, req.DeviceResource.Attributes))
		oid := req.Attributes["oid"]
		val, err := (params[i]).Int32Value()
		if err != nil {
			return err
		}
		commands = append(commands, NewSetDeviceCommand(oid, int(val)))
	}

	port := uint16(Port)
	if Port == 0 {
		port = DEFAULT_PORT
	}

	for _, command := range commands {

		var commandArray = []DeviceCommand{command}

		mu.Lock()
		if client == nil {
			client = NewSNMPClient(Address, port)
		}

		_, err2 := client.SetValues(commandArray)
		mu.Unlock()

		if err2 != nil {
			s.lc.Error(fmt.Sprintf("SNMPDriver.HandleWriteCommands; %s", err2))
			return err2
		}
	}
	return nil

}

// Stop the protocol-specific DS code to shutdown gracefully, or
// if the force parameter is 'true', immediately. The driver is responsible
// for closing any in-use channels, including the channel used to send async
// readings (if supported).
func (s *SNMPDriver) Stop(force bool) error {
	s.lc.Debug(fmt.Sprintf("SNMPDriver.Stop called: force=%v", force))
	client.Disconnect()
	return nil
}

func (d *SNMPDriver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	d.lc.Debug(fmt.Sprintf("a new Device is added: %s", deviceName))
	return nil
}

func (d *SNMPDriver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	d.lc.Debug(fmt.Sprintf("Device %s is updated", deviceName))
	return nil
}

func (d *SNMPDriver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	d.lc.Debug(fmt.Sprintf("Device %s is removed", deviceName))
	return nil
}
