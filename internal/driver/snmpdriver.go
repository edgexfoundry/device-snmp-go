// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 Dell Technologies
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"strconv"
	"sync"

	dsModels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
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
		oid, ok := req.Attributes["oid"].(string)
		if !ok {
			return nil, fmt.Errorf("oid attributes of %s must be string type", req.DeviceResourceName)
		}
		s.lc.Debugf("SNMPDriver.HandleReadCommand: device: %s operation: %v attributes: %v", Name, req.DeviceResourceName, req.Attributes)

		commands = append(commands, NewGetDeviceCommand(oid))
	}

	port, err := strconv.ParseUint(Port, 10, 64)
	if err != nil {
		s.lc.Errorf("SNMPDriver.HandleReadCommands; %s", err)
		return
	}

	mu.Lock()
	if client == nil {
		client = NewSNMPClient(Address, uint16(port))
	}

	// this can be any type either string or int at this point
	vals, err := client.GetValues(commands)

	mu.Unlock()

	if err != nil {
		s.lc.Errorf("SNMPDriver.HandleReadCommands; %s", err)
		return
	}

	for i, val := range vals {
		// switch here on type
		//s.lc.Debug(fmt.Sprintf("SNMPDriver.HandleReadCommands; value of %v is: %d", reqs[i].RO, val))

		switch safeVal := val.(type) {
		case string:
			cv, _ := dsModels.NewCommandValue(reqs[i].DeviceResourceName, common.ValueTypeString, safeVal)
			res = append(res, cv)
		case int64:
			cv, _ := dsModels.NewCommandValue(reqs[i].DeviceResourceName, common.ValueTypeInt64, safeVal)
			res = append(res, cv)
		case int32:
			cv, _ := dsModels.NewCommandValue(reqs[i].DeviceResourceName, common.ValueTypeInt32, safeVal)
			res = append(res, cv)
		case int16:
			cv, _ := dsModels.NewCommandValue(reqs[i].DeviceResourceName, common.ValueTypeInt16, safeVal)
			res = append(res, cv)
		case int8:
			cv, _ := dsModels.NewCommandValue(reqs[i].DeviceResourceName, common.ValueTypeInt8, safeVal)
			res = append(res, cv)
		case uint64:
			cv, _ := dsModels.NewCommandValue(reqs[i].DeviceResourceName, common.ValueTypeUint64, safeVal)
			res = append(res, cv)
		case uint32:
			cv, _ := dsModels.NewCommandValue(reqs[i].DeviceResourceName, common.ValueTypeUint32, safeVal)
			res = append(res, cv)
		case uint16:
			cv, _ := dsModels.NewCommandValue(reqs[i].DeviceResourceName, common.ValueTypeUint16, safeVal)
			res = append(res, cv)
		case uint8:
			cv, _ := dsModels.NewCommandValue(reqs[i].DeviceResourceName, common.ValueTypeUint8, safeVal)
			res = append(res, cv)
		case uint:
			trueVal, ok := val.(uint64) // Alt. non panicking version
			if ok {
				cv, _ := dsModels.NewCommandValue(reqs[i].DeviceResourceName, common.ValueTypeUint64, trueVal)
				res = append(res, cv)
			} else {
				s.lc.Errorf("SNMPDriver.HandleWriteCommands; %s", err)
			}
		default:
			trueVal, ok := val.(int) // Alt. non panicking version
			if ok {
				cv, _ := dsModels.NewCommandValue(reqs[i].DeviceResourceName, common.ValueTypeInt32, int32(trueVal))
				res = append(res, cv)
			} else {
				s.lc.Errorf("SNMPDriver.HandleWriteCommands; %s", err)
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
		oid, ok := req.Attributes["oid"].(string)
		if !ok {
			return fmt.Errorf("oid attributes of %s must be string type", req.DeviceResourceName)
		}
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
	s.lc.Debugf("SNMPDriver.Stop called: force=%v", force)
	client.Disconnect()
	return nil
}

func (s *SNMPDriver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	s.lc.Debugf("a new Device is added: %s", deviceName)
	return nil
}

func (s *SNMPDriver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
	s.lc.Debugf("Device %s is updated", deviceName)
	return nil
}

func (s *SNMPDriver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	s.lc.Debugf("Device %s is removed", deviceName)
	return nil
}
