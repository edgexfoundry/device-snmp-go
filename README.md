# edgexfoundry/device-snmp-go
[![Build Status](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/device-snmp-go/job/main/badge/icon)](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/device-snmp-go/job/main/) [![Code Coverage](https://codecov.io/gh/edgexfoundry/device-snmp-go/branch/main/graph/badge.svg?token=cQs6ERHT74)](https://codecov.io/gh/edgexfoundry/device-snmp-go) [![Go Report Card](https://goreportcard.com/badge/github.com/edgexfoundry/device-snmp-go)](https://goreportcard.com/report/github.com/edgexfoundry/device-snmp-go) [![GitHub Latest Dev Tag)](https://img.shields.io/github/v/tag/edgexfoundry/device-snmp-go?include_prereleases&sort=semver&label=latest-dev)](https://github.com/edgexfoundry/device-snmp-go/tags) ![GitHub Latest Stable Tag)](https://img.shields.io/github/v/tag/edgexfoundry/device-snmp-go?sort=semver&label=latest-stable) [![GitHub License](https://img.shields.io/github/license/edgexfoundry/device-snmp-go)](https://choosealicense.com/licenses/apache-2.0/) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/edgexfoundry/device-snmp-go) [![GitHub Pull Requests](https://img.shields.io/github/issues-pr-raw/edgexfoundry/device-snmp-go)](https://github.com/edgexfoundry/device-snmp-go/pulls) [![GitHub Contributors](https://img.shields.io/github/contributors/edgexfoundry/device-snmp-go)](https://github.com/edgexfoundry/device-snmp-go/contributors) [![GitHub Committers](https://img.shields.io/badge/team-committers-green)](https://github.com/orgs/edgexfoundry/teams/device-snmp-go-committers/members) [![GitHub Commit Activity](https://img.shields.io/github/commit-activity/m/edgexfoundry/device-snmp-go)](https://github.com/edgexfoundry/device-snmp-go/commits)


#### Overview:
A generic SNMP device service. Included is a sample for a network switch device driver. Additionally, a Patlite profile is included in the profiles director. 

This device service facilitates via the configuration, accessing the TrendNet TEG-082WS 10 port switch.
You can send and receive event readings from the network switch.  Currently this device service is tailored specifically to the TrendNET TEG-082WS switch but refactoring could take place to abstract the device type and provde a myrida of device profiles. You would need to interogate the network switch and then determine the appropriate device profile.   This particular device service uses the SNMP (Simple Netowrk Management Protocol).
SNMP communication with the switch permits the device service to create events based on the Network Switch data. For further information on the Simple Network Mangement Protocol [`SNMP Wiki`][ExWk]

#### TODOS
> Async callbacks not supported
> SNMP Traps are currently not implemented but will be completed at a future date. 
> Network Switch Type abstraction


### Pre-Requisites

device-snmp-go requires [Go](https://golang.org/dl/) language. Download the current release if you don't have it already.
Have access to the  [`EdgeX Foundry`](https://github.com/edgexfoundry) project github site.  

### EdgeX-Go (The EdgeX Platform and Sevices)
This service requires the EdgeX-Go platform and serveral of its services to be up and running if you want to test the device service full integration. 
However, you can run this service as a standalone with the following minium services up and running either a) natively (.exe) or b) as Docker image containers.

Detailed instructions on how to use the various Edgex-Foundry service(s) are linked below.  

| Service | README |
| ------ | ------ |
| core-data | [core/core-data/README.md][ExCd] |
| core-command | [core/core-command/README.md][ExCc] |
| core-metadata | [core/core-metadata/README.md][ExCm] |

### EdgeX Foundry Images and Containers
If you choose to run the required services or all services as Docker containers you will need to use the Delhi version 0.7.1 of the Edgex-Go Services.

Download from GitHub the EdgeXFoundry [`developer-scripts`](https://github.com/edgexfoundry/developer-scripts) project.  I download all of my EdgeX Foundry repositories into a parent 'edgexfoundry' directory. The following docker commands assume you have that parent directory in place.  

To download the required Delhi versions of the Edgex-Go services perform the following from the console.
```sh
cd .../edgexfoundry/developer-scripts/compose-files/
docker-compose up -d 
```
Now to stop all of the services you and do the following:  
```sh
cd .../edgexfoundry/developer-scripts/compose-files/
docker-compose down  
```
Now to stop all of the services and remove the docker volumes do the following:
```sh
cd .../edgexfoundry/developer-scripts/compose-files/
docker-compose down -v 
```
You would want to remove the docker image volumes for anything you persisted.  The volumes contain all the data you have published to your edgex-go-mongo as an example. 
As a general practice I routinely remove my docker volumes.   For more information on docker commands see ref [`Docker Commands`](https://docs.docker.com/engine/reference/commandline/docker/)
> Note:  docker-compose.yml is going to be the latest release version of the edgex-foundry project.  
> For future versions please use the docker-compose-delhi-0.7.1.yml file


#### GET calls (REST)
For device service get calls we will use our configured static IP address for our network card.  You can find this by executing the following:
osx.  Note you will need to ensure you select your working machine's ethernet static and not any other devices you have on your local network.
I have included the SNMP MIB table OID entry(s) so you can see what OID's the device service is using. 
```sh
arp -a
cheese.attlocal.net (192.168.1.69) [ethernet] ...
...
```
This is the IP address you will substititue for localhost in the following commands. 
Verify the deployment by navigating to your server address in your preferred browser or [`Postman`](https://www.getpostman.com/) REST test harness.
Or you can use the curl command from your console if your on linux or OSX.  If you choose to use Postman just use the address less the curl command. 

```sh
$ curl http://localhost:59993/api/v2/ping
```

#### Reboot 
The current state of the Network Switch reboot
##### (SNMP OID)  `{ oid: "1.3.6.1.4.1.28866.2.37.24.5.1.0", community: "private"  }`

```sh
$ curl http://localhost:59993/api/v2/trendnet01/Reboot
```

#### UpTime 
The switch uptime since last reboot in TimeTicks 
##### (SNMP OID) `{ oid: "1.3.6.1.4.1.28866.2.37.16.1.1.0", community: "private"  }`

```sh
$ curl http://localhost:59993/api/v2/trendnet01/Uptime
```
#### Firmware 
The switch firmware version 
##### (SNMP OID) ` { oid: "1.3.6.1.4.1.28866.2.37.16.1.2.0", community: "private"  } `

```sh
$ curl http://localhost:59993/api/v2/trendnet01/Firmware
```
#### MACAddress
The switch MAC Address 
##### (SNMP OID) `{ oid: "1.3.6.1.4.1.28866.2.37.16.2.1.0", community: "private"  } `

```sh
$ curl http://localhost:59993/api/v2/trendnet01/MacAddress
```
#### IPV4Address
The switch IPV4 Address 
##### (SNMP OID) `{ oid: "1.3.6.1.4.1.28866.2.37.16.1.2.0", community: "private"  } `

```sh
$ curl http://localhost:59993/api/v2/trendnet01/IPV4Address
```
#### IPV4SubnetMask
The switch Subnet Mask 
##### (SNMP OID) `{ oid: "1.3.6.1.4.1.28866.2.37.16.3.3.0", community: "private"  } `

```sh
$ curl http://localhost:59993/api/v2/trendnet01/IPV4SubnetMask
```
#### IPV4GatewayAddress
The switch IPV4 Gateway Address
##### (SNMP OID) `{ oid: "1.3.6.1.4.1.28866.2.37.16.3.4.0", community: "private"  } `

```sh
$ curl http://localhost:59993/api/v2/trendnet01/IPV4GatewayAddress
```
#### PUT calls (REST)
For the set commands we presently just have one.  I have included the SNMP MIB OID table entry(s) so you can see what OID's the device service is using. 

#### Reboot
Restarts the switch
##### (SNMP OID) ` { oid: "1.3.6.1.4.1.28866.2.37.24.5.1.0", community: "private"  } `

| Parameter | Value | Description |
| ------ | ------ | ------ | 
| RebootControlState | 1 | Restarts the switch |
| RebootControlState | 2 | Factory reset and restarts switch |
| RebootControlState | 3 | Factory reset, restarts switch with no IP Address |

```sh
$ curl -H 'Content-Type: application/json' -X PUT -d '{"RebootControlState":"1"}' http://localhost:59993/api/v2/Reboot
```

### Configuration

This service is driven by a `configuration.toml` [TOML-based](https://en.wikipedia.org/wiki/TOML) file that resides
in the directory the application is executed from. See [cmd/res/configuration.toml](cmd/res/configuration.toml) for default configuration.
### References
> Website: [Dillinger.io](https://dillinger.io) "A great README.md file editor"


[ExCd]: <https://github.com/edgexfoundry/edgex-go/tree/master/internal/core/data/README.md>
[ExCc]: <https://github.com/edgexfoundry/edgex-go/tree/master/internal/core/command/README.md>
[ExCm]: <https://github.com/edgexfoundry/edgex-go/tree/master/internal/core/metadata/README.md>
[ExWk]: <https://en.wikipedia.org/wiki/Simple_Network_Management_Protocol>