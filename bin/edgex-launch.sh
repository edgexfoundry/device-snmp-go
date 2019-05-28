#!/usr/bin/env bash
#
# Copyright (c) 2019
# Dell/EMC
#
# SPDX-License-Identifier: Apache-2.0
#

###
# Launches all EdgeX Go binaries (must be previously built).
#
# Expects that Consul and MongoDB are already installed and running.
#
###

DIR=$PWD
CMD=../cmd

function cleanup {
	pkill device-snmp-go
}

cd $CMD
exec -a device-snmp-go ./device-snmp-go &
cd $DIR


trap cleanup EXIT

while : ; do sleep 1 ; done