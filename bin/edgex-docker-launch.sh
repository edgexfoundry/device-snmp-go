#!/usr/bin/env sh

# Copyright (c) 2018
# Cavium
#
# SPDX-License-Identifier: Apache-2.0
#

# Assumes you have the EdgeX Foundry services up and running.
# SEE BELOW to
# Start EdgeX Foundry services in right order, as described:
# https://docs.edgexfoundry.org/Ch-GettingStartedUsers.html

# EDGEX_COMPOSE_FILE overrides the default compose file.
# EDGEX_SERVICES lists the service to start
#
# E.g. Start everything but metadata and data, use Redis for Core Services and a local compose file
# EDGEX_SERVICES="logging command export-client export-distro notifications" \
# EDGEX_CORE_DB=redis EDGEX_COMPOSE_FILE=docker/local-docker-compose.yml \
# bin/edgex-docker-launch.sh


COMPOSE_FILE=$EDGEX_COMPOSE_FILE

echo "Starting device-snmp-switch-go"
docker-compose -f $COMPOSE_FILE up -d device-snmp-switch-go

echo "Sleeping before launching remaining services"
sleep 5

defaultServices="device-snmp-switch-go"
if [ -z ${EDGEX_SERVICES} ]; then
  deps=
  services=${defaultServices}
else
  deps=--no-deps
  services=${EDGEX_SERVICES}
fi

for s in ${services}; do
    echo Starting ${s}
    docker-compose -f $COMPOSE_FILE up -d ${deps} $s
done