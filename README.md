# edgexfoundry/device-snmp-go

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
$ curl http://localhost:49993/api/v1/ping
```

#### Reboot 
The current state of the Network Switch reboot
##### (SNMP OID)  `{ oid: "1.3.6.1.4.1.28866.2.37.24.5.1.0", community: "private"  }`

```sh
$ curl http://localhost:49993/api/v1/trendnet01/Reboot
```

#### UpTime 
The switch uptime since last reboot in TimeTicks 
##### (SNMP OID) `{ oid: "1.3.6.1.4.1.28866.2.37.16.1.1.0", community: "private"  }`

```sh
$ curl http://localhost:49993/api/v1/trendnet01/Uptime
```
#### Firmware 
The switch firmware version 
##### (SNMP OID) ` { oid: "1.3.6.1.4.1.28866.2.37.16.1.2.0", community: "private"  } `

```sh
$ curl http://localhost:49993/api/v1/trendnet01/Firmware
```
#### MACAddress
The switch MAC Address 
##### (SNMP OID) `{ oid: "1.3.6.1.4.1.28866.2.37.16.2.1.0", community: "private"  } `

```sh
$ curl http://localhost:49993/api/v1/trendnet01/MacAddress
```
#### IPV4Address
The switch IPV4 Address 
##### (SNMP OID) `{ oid: "1.3.6.1.4.1.28866.2.37.16.1.2.0", community: "private"  } `

```sh
$ curl http://localhost:49993/api/v1/trendnet01/IPV4Address
```
#### IPV4SubnetMask
The switch Subnet Mask 
##### (SNMP OID) `{ oid: "1.3.6.1.4.1.28866.2.37.16.3.3.0", community: "private"  } `

```sh
$ curl http://localhost:49993/api/v1/trendnet01/IPV4SubnetMask
```
#### IPV4GatewayAddress
The switch IPV4 Gateway Address
##### (SNMP OID) `{ oid: "1.3.6.1.4.1.28866.2.37.16.3.4.0", community: "private"  } `

```sh
$ curl http://localhost:49993/api/v1/trendnet01/IPV4GatewayAddress
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
$ curl -H 'Content-Type: application/json' -X PUT -d '{"RebootControlState":"1"}' http://localhost:49993/api/v1/Reboot
```

### Configuration

This service is driven by a `configuration.toml` [TOML-based](https://en.wikipedia.org/wiki/TOML) file that resides
    in the directory the application is executed from.  Configuration has several required key/value pairs:
    
| Parent | Entry | Description |
| ----- | ----- | ----- |
|``[Writable]``|| configuration entry for the `Writeable`. Sections which can be overwritten via external or internal configuration. | 
|``[Service]``|| configuration entry for the `Service` which will be running | 
||`ServerHost`| a string, this defines the domain or ip address the application should bind to for handling inbound dashboard requests. It's also the address used when assembling the callback uri provided to EdgeX's export-distro service to deliver events.  If you follow the [Start Up](#start-up) directions, it shouldn't be `localhost` since it needs to be a bindable address in your local host environment that can be reached from inside the edgex-export-distro container.  On my development machine, I used my host's IP address and updated my anti-virus's software firewall to allow the `ServerPort` through for all addresses. |
||`ServerPort`| an int, this defines the TCP/IP port the application should bind to for handling inbound dashboard requests.|
||`ConnectionRetries`| an int, this defines the maximum retry attempts to connect to this service.|
||`Labels`|an array of items defining the service.|
||`OpenMsg`|a string, this is the opening message displayed when launching the service from the `edgex-go-sdk` bootstrap process.|
||`ReadMaxLimit`|an int, this defines the maximum read size for characters from the service.|
||`Timeout`|an int, this defines the server timeout max duration in milliseconds. Example: 5000|
||`EnableAsyncReadings`|a `true` or `false` entry. This enables or disables asynchronous readings from the service.|
||`AsyncBufferSize`|an int, this defines the maximum asynchronous buffer size.|
|``[Registry]``||configuration entry for the `Registry Service`|
|| `Host`| a string for the host ip address|
|| `Port`| an int for the host port|
|| `CheckInterval`| an int representing the preriodicity for the registry to check|
|| `FailLimit`| an int representing the max connection failure attempts|
|| `FailWaitTime`| an int representing the duration in seconds between subsequent fail requests |
|`[Clients]`||configuration parent for `N number of Clients`|
||` [Clients.XXX]`| generic configuration type which uses the Client struct. You must have each __`Clients.Name`__ unique.  Example would be `[Clients.Data]` this will provide a unique client configuration. Below are the lists of interface or struct fields which are common to the `[Clients]` type.|
||`Name`| a string representing the client name|
||`Protocol`| a string representing the protocol |
||`Host`| a string representing the host IP Address |
||`Port`| an int representing the port for the host |
||`Timeout`| an int representing the connection timeout |
|`[Device]`||configuration entry for a `Device Service`|
||`DataTransform`|a `true` or `false` should you transform the data|
||`InitCmd`|a string representing the intial command to execute|
||`InitCmdArgs`|a string representing the intital command arguments|
||`MaxCmdOps`|an int representing the maximum number of command opertions concurrently|
||`MaxCmdValueLn`|an int representing the maximum character limit for a command string|
||`RemoveCmd`|a string respresenting the command you want to remove|
||`RemoveCmdArgs`|a string representing the command arguments you want to remove|
||`ProfilesDir`|a string representing the directory relative to the running application. example: "./res"|
|`[Logging]`||configuration representing the `Logging Service`|
||`EnableRemote`|a `true` or `false` remote logging|
||`File`|a string representing the file path for the local logging file. example: "./device-snmp-device.log"|
||`Level`|a string representing the Log Levels. Example: `DEBUG`,`WARN`, `ERROR`, or `INFO`|
|`[[DeviceList]]`||configuration entry representing a list of device services.|
||`Name`| a string representing the name of the device to be configured. This is the actual hardware device name.|
||`Profile`| a string representing the `Profile` name to be used in edgex-core-command and edgex-core-metadata. This is the profile for your given device which will be listed via the core-metadata service REST API.|
||`[DeviceList.Protocols`|configuration representing the protocol for a given Device|
||`Address`| a string representing the IP address for the device.  This should be relative to your network which you are using.|
||`Port`|an int representing the `Port` for your device|
||`Protocol`|a string representing the `Protocol` used to communicate with the configured device. Example: `HTTP`, or `TCP`.|
|`[[DeviceList.AutoEvents]]`||configuration representing the scheduling for a periodic interval in time. Note: This can occur more then once.  So you could have `N number` of complete `[[AutoEvents]]` defined in your configuration, for more then one time interval defined. |
||`Frequency`|a string representing the time interval in CRON format. see [CRON FORMAT](http://www.nncron.ru/help/EN/working/cron-format.htm)|
||`Resource`| string representing the `unique` name for a given device resource.|
||`Port`|an int representing the `Port` for your device|



### References
> Website: [Dillinger.io](https://dillinger.io) "A great README.md file editor"


[ExCd]: <https://github.com/edgexfoundry/edgex-go/tree/master/internal/core/data/README.md>
[ExCc]: <https://github.com/edgexfoundry/edgex-go/tree/master/internal/core/command/README.md>
[ExCm]: <https://github.com/edgexfoundry/edgex-go/tree/master/internal/core/metadata/README.md>
[ExWk]: <https://en.wikipedia.org/wiki/Simple_Network_Management_Protocol>