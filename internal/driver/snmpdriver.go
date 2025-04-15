// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 Dell Technologies
// Copyright (C) 2020-2021,2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/edgexfoundry/device-sdk-go/v4/pkg/interfaces"
	dsModels "github.com/edgexfoundry/device-sdk-go/v4/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

type SNMPDriver struct {
	lc      logger.LoggingClient
	asyncCh chan<- *dsModels.AsyncValues
}

var client *SNMPClient

// Used to avoid get/set at the same time. If this happens simultaneously, state
// of the device can get out of sync with command actuation result
var mu sync.Mutex

const (
	TcpProtocol     = "TCP"
	AddressProperty = "Address"
	PortProperty    = "Port"
)

// DisconnectDevice handles protocol-specific cleanup when a device
// is removed.
func (s *SNMPDriver) DisconnectDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
	s.lc.Warn("Driver's DisconnectDevice function didn't implement")
	return nil
}

// Initialize performs protocol-specific initialization for the device
// service.
func (s *SNMPDriver) Initialize(sdk interfaces.DeviceServiceSDK) error {
	s.lc = sdk.LoggingClient()
	s.asyncCh = sdk.AsyncValuesChannel()
	return nil
}

func (s *SNMPDriver) Start() error {
	return nil
}

// HandleReadCommands triggers a protocol Read operation for the specified device.
func (s *SNMPDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	res = make([]*dsModels.CommandValue, len(reqs))

	var TCP = protocols[TcpProtocol]

	Name := deviceName
	Address := TCP[AddressProperty].(string)
	Port := TCP[PortProperty].(string)

	//s.lc.Debug(fmt.Sprintf("SimpleDriver.HandleReadCommands: protocols: %v operation: %v attributes: %v", protocols, reqs[0].RO.Operation, reqs[0].DeviceResource.Attributes))

	s.lc.Debugf("Address:%s, Port:%s", Address, Port)
	port, err := strconv.ParseUint(Port, 10, 64)
	if err != nil {
		s.lc.Errorf("SNMPDriver.HandleReadCommands; %s", err)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	if client == nil {
		client = NewSNMPClient(Address, uint16(port))
	}
	var failedCount = 0
	for i, req := range reqs {
		commands := make([]DeviceCommand, 1)
		oid, ok := req.Attributes["oid"].(string)
		if !ok {
			return nil, fmt.Errorf("oid attributes of %s must be string type", req.DeviceResourceName)
		}
		s.lc.Debugf("SNMPDriver.HandleReadCommand: device: %s operation: %v attributes: %v", Name, req.DeviceResourceName, req.Attributes)

		commands[0] = NewGetDeviceCommand(oid)
		val, err := client.GetValues(commands)
		if err != nil {
			failedCount++
			s.lc.Warnf("SNMPDriver.HandleReadCommands; err:%v", err)
			res[i] = nil
			continue
		}
		s.lc.Debugf("oid %v get value %v", oid, val)
		res[i], err = s.switchType(req, val[0])
		if err != nil {
			failedCount++
			s.lc.Warnf("switch type failed; err:%v", err)
			res[i] = nil
			continue
		}
	}
	if failedCount == len(reqs) {
		s.lc.Errorf("handleReadCommands: Handle read commandRequest %+v failed", reqs)
		return res, fmt.Errorf("handleReadCommands: Handle read commandRequest %+v failed", reqs)
	}
	s.lc.Debugf("get result %+v", res)
	return
}

func (s *SNMPDriver) switchType(req dsModels.CommandRequest, val interface{}) (result *dsModels.CommandValue, err error) {
	// switch here on type
	//s.lc.Debug(fmt.Sprintf("SNMPDriver.HandleReadCommands; value of %v is: %d", reqs[i].RO, val))
	switch safeVal := val.(type) {
	case string:
		result, err = dsModels.NewCommandValue(req.DeviceResourceName, common.ValueTypeString, safeVal)
	case int64:
		result, err = dsModels.NewCommandValue(req.DeviceResourceName, common.ValueTypeInt64, safeVal)
	case int32:
		result, err = dsModels.NewCommandValue(req.DeviceResourceName, common.ValueTypeInt32, safeVal)
	case int16:
		result, err = dsModels.NewCommandValue(req.DeviceResourceName, common.ValueTypeInt16, safeVal)
	case int8:
		result, err = dsModels.NewCommandValue(req.DeviceResourceName, common.ValueTypeInt8, safeVal)
	case uint64:
		result, err = dsModels.NewCommandValue(req.DeviceResourceName, common.ValueTypeUint64, safeVal)
	case uint32:
		result, err = dsModels.NewCommandValue(req.DeviceResourceName, common.ValueTypeUint32, safeVal)
	case uint16:
		result, err = dsModels.NewCommandValue(req.DeviceResourceName, common.ValueTypeUint16, safeVal)
	case uint8:
		result, err = dsModels.NewCommandValue(req.DeviceResourceName, common.ValueTypeUint8, safeVal)
	case uint:
		trueVal, ok := val.(uint64) // Alt. non panicking version
		if ok {
			result, err = dsModels.NewCommandValue(req.DeviceResourceName, common.ValueTypeUint64, trueVal)
		} else {
			err = fmt.Errorf("cannot convert variable to uint64 type")
		}
	default:
		trueVal, ok := val.(int) // Alt. non panicking version
		if ok {
			result, err = dsModels.NewCommandValue(req.DeviceResourceName, common.ValueTypeInt32, int32(trueVal))
		} else {
			err = fmt.Errorf("cannot convert variable to int type")
		}
	}
	if err != nil {
		s.lc.Errorf("SNMPDriver.HandleReadCommands; err:%s; req:%+v", err, req)
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

func (s *SNMPDriver) Discover() error {
	return fmt.Errorf("driver's Discover function isn't implemented")
}

func (s *SNMPDriver) ValidateDevice(device models.Device) error {
	protocol, ok := device.Protocols[TcpProtocol]
	if !ok {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("protocol '%s' does not exist", TcpProtocol), nil)
	}
	address := ""
	if v, ok := protocol[AddressProperty]; ok {
		address = fmt.Sprintf("%v", v)
	}
	if address == "" {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("protocol properties '%s' does not exist", AddressProperty), nil)
	}
	port := ""
	if v, ok := protocol[PortProperty]; ok {
		port = fmt.Sprintf("%v", v)
		_, err := strconv.ParseUint(port, 10, 64)
		if err != nil {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid protocol properties '%s': %v", PortProperty, port), err)
		}
	}
	if port == "" {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("protocol properties '%s' does not exist", PortProperty), nil)
	}
	return nil
}
