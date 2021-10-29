// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2021 Canonical Ltd
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 * SPDX-License-Identifier: Apache-2.0'
 */

package hooks

// ConfToEnv defines mappings from snap config keys to EdgeX environment variable
// names that are used to override individual device-snmp's [Driver]  configuration
// values via a .env file read by the snap service wrapper.
//
// The syntax to set a configuration key is:
//
// env.<section>.<keyname>
//
var ConfToEnv = map[string]string{
	// [snmpBrokerInfo]
	"snmpbrokerinfo.schema":     "snmpBROKERINFO_SCHEMA",
	"snmpbrokerinfo.host":       "snmpBROKERINFO_HOST",
	"snmpbrokerinfo.port":       "snmpBROKERINFO_PORT",
	"snmpbrokerinfo.qos":        "snmpBROKERINFO_QOS",
	"snmpbrokerinfo.keep-alive": "snmpBROKERINFO_KEEPALIVE",
	"snmpbrokerinfo.client-id":  "snmpBROKERINFO_CLIENTID",

	"snmpbrokerinfo.credentials-retry-time":  "snmpBROKERINFO_CREDENTIALSRETRYTIME",
	"snmpbrokerinfo.credentials-retry-wait":  "snmpBROKERINFO_CREDENTIALSRETRYWAIT",
	"snmpbrokerinfo.conn-establishing-retry": "snmpBROKERINFO_CONNESTABLISHINGRETRY",
	"snmpbrokerinfo.conn-retry-wait-time":    "snmpBROKERINFO_CONNRETRYWAITTIME",

	"snmpbrokerinfo.auth-mode":        "snmpBROKERINFO_AUTHMODE",
	"snmpbrokerinfo.credentials-path": "snmpBROKERINFO_CREDENTIALSPATH",

	"snmpbrokerinfo.incoming-topic": "snmpBROKERINFO_INCOMINGTOPIC",
	"snmpbrokerinfo.response-topic": "snmpBROKERINFO_RESPONSETOPIC",

	// [Device]
	"device.update-last-connected": "DEVICE_UPDATELASTCONNECTED",
	"device.use-message-bus":       "DEVICE_USEMESSAGEBUS",
}
