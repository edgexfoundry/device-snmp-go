
<a name="EdgeX SNMP Device Service (found in device-snmp-go) Changelog"></a>
## EdgeX SNMP Device Service
[Github repository](https://github.com/edgexfoundry/device-snmp-go)

## [v2.0.0] Ireland - 2021-06-30  (Not Compatible with 1.x releases)
### Change Logs for EdgeX Dependencies
- [device-sdk-go](https://github.com/edgexfoundry/device-sdk-go/blob/v2.0.0/CHANGELOG.md)
- [go-mod-core-contracts](https://github.com/edgexfoundry/go-mod-core-contracts/blob/v2.0.0/CHANGELOG.md)

### Features ✨
- Enable using MessageBus as the default ([#37e4736](https://github.com/edgexfoundry/device-snmp-go/commits/37e4736))
- Remove Logging configuration ([#9643cb4](https://github.com/edgexfoundry/device-snmp-go/commits/9643cb4))
### Bug Fixes 🐛
### Code Refactoring ♻
- remove unimplemented InitCmd/RemoveCmd configuration ([#07622c1](https://github.com/edgexfoundry/device-snmp-go/commits/07622c1))
- Change PublishTopicPrefix value to be 'edgex/events/device' ([#a95e588](https://github.com/edgexfoundry/device-snmp-go/commits/a95e588))
- Update configuration for change to common ServiceInfo struct ([#47d5e2d](https://github.com/edgexfoundry/device-snmp-go/commits/47d5e2d))
    ```
    BREAKING CHANGE:
    Service configuration has changed
    ```
- Update to assign and uses new Port Assignments ([#65e5e53](https://github.com/edgexfoundry/device-snmp-go/commits/65e5e53))
    ```
    BREAKING CHANGE:
    Device SNMP default port number has changed to 59993
    ```
- update for new service key name ([#1279fd2](https://github.com/edgexfoundry/device-snmp-go/commits/1279fd2))
- use v2 device-sdk ([#4f41517](https://github.com/edgexfoundry/device-snmp-go/commits/4f41517))
### Documentation 📖
- Add badges to readme ([#e4f17ee](https://github.com/edgexfoundry/device-snmp-go/commits/e4f17ee))
### Build 👷
- update build files for v2 ([#6247689](https://github.com/edgexfoundry/device-snmp-go/commits/6247689))
- update Dockerfiles to use go 1.16 ([#029de77](https://github.com/edgexfoundry/device-snmp-go/commits/029de77))
- update go.mod to go 1.16 ([#1535d2d](https://github.com/edgexfoundry/device-snmp-go/commits/1535d2d))
### Continuous Integration 🔄
- update local docker image names ([#639fbe6](https://github.com/edgexfoundry/device-snmp-go/commits/639fbe6))

<a name="v1.3.1"></a>
## [v1.3.1] - 2021-02-02
### Bug Fixes 🐛
- updated all example profiles to Hanoi profile structure ([#28132d5](https://github.com/edgexfoundry/device-snmp-go/commits/28132d5))
### Build 👷
- **deps:** bump github.com/edgexfoundry/device-sdk-go ([#fc9faa1](https://github.com/edgexfoundry/device-snmp-go/commits/fc9faa1))
### Continuous Integration 🔄
- add semantic.yml for commit linting, update PR template to latest ([#0f8f4ec](https://github.com/edgexfoundry/device-snmp-go/commits/0f8f4ec))
- standardize dockerfiles ([#12d1790](https://github.com/edgexfoundry/device-snmp-go/commits/12d1790))

<a name="v1.3.0"></a>
## [v1.3.0] - 2020-11-18
### Code Refactoring ♻
- Upgrade SDK to v1.2.4-dev.34 ([#892246b](https://github.com/edgexfoundry/device-snmp-go/commits/892246b))
- update dockerfile to appropriately use ENTRYPOINT and CMD, closes[#57](https://github.com/edgexfoundry/device-snmp-go/issues/57) ([#4de3236](https://github.com/edgexfoundry/device-snmp-go/commits/4de3236))
### Build 👷
- Upgrade to Go1.15 ([#2a9ba1b](https://github.com/edgexfoundry/device-snmp-go/commits/2a9ba1b))
- **all:** Enable use of DependaBot to maintain Go dependencies ([#13a45a7](https://github.com/edgexfoundry/device-snmp-go/commits/13a45a7))

<a name="v1.2.2"></a>
## [v1.2.2] - 2020-08-19
### Code Refactoring ♻
- upgrade SDK to v1.2.0 ([#2df6a38](https://github.com/edgexfoundry/device-snmp-go/commits/2df6a38))
### Documentation 📖
- Add standard PR template ([#a0c73d4](https://github.com/edgexfoundry/device-snmp-go/commits/a0c73d4))
### Build 👷
- Update relevant files in device-snmp-go for Go 1.13. ([#02f9db0](https://github.com/edgexfoundry/device-snmp-go/commits/02f9db0))

<a name="v1.1.1"></a>
## [v1.1.1] - 2019-12-13

<a name="v1.1.0"></a>
## [v1.1.0] - 2019-11-18
### Pipeline
- bring in updated Jenkinsfile from edinburgh ([#1565ae0](https://github.com/edgexfoundry/device-snmp-go/commits/1565ae0))

<a name="1.0.0"></a>
## [1.0.0] - 2019-06-24
### Pipeline
- bump pre instead of patch ([#bc723af](https://github.com/edgexfoundry/device-snmp-go/commits/bc723af))

