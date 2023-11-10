
<a name="EdgeX SNMP Device Service (found in device-snmp-go) Changelog"></a>
## EdgeX SNMP Device Service
[Github repository](https://github.com/edgexfoundry/device-snmp-go)

### Change Logs for EdgeX Dependencies
- [device-sdk-go](https://github.com/edgexfoundry/device-sdk-go/blob/main/CHANGELOG.md)
- [go-mod-core-contracts](https://github.com/edgexfoundry/go-mod-core-contracts/blob/main/CHANGELOG.md)
- [go-mod-bootstrap](https://github.com/edgexfoundry/go-mod-bootstrap/blob/main/CHANGELOG.md)  (indirect dependency)
- [go-mod-messaging](https://github.com/edgexfoundry/go-mod-messaging/blob/main/CHANGELOG.md) (indirect dependency)
- [go-mod-registry](https://github.com/edgexfoundry/go-mod-registry/blob/main/CHANGELOG.md)  (indirect dependency)
- [go-mod-secrets](https://github.com/edgexfoundry/go-mod-secrets/blob/main/CHANGELOG.md) (indirect dependency)
- [go-mod-configuration](https://github.com/edgexfoundry/go-mod-configuration/blob/main/CHANGELOG.md) (indirect dependency)


## [v3.1.0] Napa - 2023-11-15 (Only compatible with the 3.x releases)


### ✨  Features

- Remove snap packaging ([#322](https://github.com/edgexfoundry/device-snmp-go/issues/322)) ([c9f5e79…](https://github.com/edgexfoundry/device-snmp-go/commit/c9f5e79f7f6757774f639370f81fc5ed0e697313))
```text

BREAKING CHANGE: Remove snap packaging ([#322](https://github.com/edgexfoundry/device-snmp-go/issues/322))

```


### ♻ Code Refactoring

- Remove obsolete comments from config file ([#323](https://github.com/edgexfoundry/device-snmp-go/issues/323)) ([e3ad67b…](https://github.com/edgexfoundry/device-snmp-go/commit/e3ad67b18069b788d3b650faffca04e548476692))
- Remove github.com/pkg/errors from Attribution.txt ([c47c4ed…](https://github.com/edgexfoundry/device-snmp-go/commit/c47c4ed6f33dd6b026481a957b2bee764cde5d0f))


### 👷 Build

- Upgrade to go-1.21, Linter1.54.2 and Alpine 3.18 ([#317](https://github.com/edgexfoundry/device-snmp-go/issues/317)) ([1e145ee…](https://github.com/edgexfoundry/device-snmp-go/commit/1e145ee1301df0b97f405486925368d7c560441a))


### 🤖 Continuous Integration

- Add automated release workflow on tag creation ([67706f8…](https://github.com/edgexfoundry/device-snmp-go/commit/67706f8f4f9b3012e3ba766cacf3a9b27c779e94))


## [v3.0.0] Minnesota - 2023-05-31 (Only compatible with the 3.x releases)

### Features ✨
- Add device validation function ([#284](https://github.com/edgexfoundry/device-snmp-go/pull/284))
    ```text
    BREAKING CHANGE: Implement `ValidateDevice` function to validate device protocol properties for core-metadata  
    ```
- Update for common config ([#255](https://github.com/edgexfoundry/device-snmp-go/pull/255))
    ```text
    BREAKING CHANGE: Configuration file is changed to remove common config settings
    ```
- Use latest SDK for MessageBus Request API ([#211](https://github.com/edgexfoundry/device-snmp-go/pull/211))
    ```text
    BREAKING CHANGE: Commands via MessageBus topic configuration are changed
    ```
- Remove ZeroMQ MessageBus capability ([#210](https://github.com/edgexfoundry/device-snmp-go/pull/210))
    ```text
    BREAKING CHANGE: ZeroMQ MessageBus capability no longer available
    ```

### Bug Fixes 🐛
- **snap:** Refactor to avoid conflicts with readonly config provider directory ([#270](https://github.com/edgexfoundry/device-snmp-go/issues/270)) ([#224900f](https://github.com/edgexfoundry/device-snmp-go/commits/224900f))

### Code Refactoring ♻
- Change configuration and devices files format to YAML ([#279](https://github.com/edgexfoundry/device-snmp-go/pull/279))
    ```text
    BREAKING CHANGE: Configuration files are now in YAML format, Default file name is now configuration.yaml
    ```
- **snap:** Update command and metadata sourcing ([#266](https://github.com/edgexfoundry/device-snmp-go/issues/266)) ([#965f04e](https://github.com/edgexfoundry/device-snmp-go/commits/965f04e))
- **snap:** Drop the support for legacy snap env options ([#213](https://github.com/edgexfoundry/device-snmp-go/issues/213))
    ```text
    BREAKING CHANGE:
    - Drop the support for legacy snap options with env. prefix
    - Upgrade edgex-snap-hooks to v3
    - Upgrade edgex-snap-testing Github action to v3
    - Add snap's Go module to dependabot
    - Other minor refactoring
    ```

### Build 👷
- Update to Go 1.20, Alpine 3.17 and linter v1.51.2 ([#06818cd](https://github.com/edgexfoundry/device-snmp-go/commits/06818cd))

## [v2.3.0] Levski - 2022-11-09 (Not Compatible with 1.x releases)

### Features ✨

- Add Service Metrics configuration ([#200](https://github.com/edgexfoundry/device-snmp-go/issues/200)) ([#073f662](https://github.com/edgexfoundry/device-snmp-go/commits/073f662))
- Add NATS configuration and build option ([#191](https://github.com/edgexfoundry/device-snmp-go/issues/191)) ([#60caeb9](https://github.com/edgexfoundry/device-snmp-go/commits/60caeb9))
- Add commanding via message configuration ([#7cd4731](https://github.com/edgexfoundry/device-snmp-go/commits/7cd4731))
- Add go-winio to attribution (new SPIFFE dependency) ([#175](https://github.com/edgexfoundry/device-snmp-go/issues/175)) ([#c4b573a](https://github.com/edgexfoundry/device-snmp-go/commits/c4b573a))
- **snap:** add config interface with unique identifier ([#197](https://github.com/edgexfoundry/device-snmp-go/issues/197)) ([#c5649dd](https://github.com/edgexfoundry/device-snmp-go/commits/c5649dd))

### Code Refactoring ♻

- **snap:** edgex-snap-hooks related upgrade ([#185](https://github.com/edgexfoundry/device-snmp-go/issues/185)) ([#a30604e](https://github.com/edgexfoundry/device-snmp-go/commits/a30604e))

### Build 👷

- Upgrade to Go 1.18 and optimize attribution script ([#182](https://github.com/edgexfoundry/device-snmp-go/issues/182)) ([#d0a0e99](https://github.com/edgexfoundry/device-snmp-go/commits/d0a0e99))

## [v2.2.0] Kamakura - 2022-05-11  (Not Compatible with 1.x releases)

### Features ✨
- Enable security hardening ([#8ca7183](https://github.com/edgexfoundry/device-snmp-go/commits/8ca7183))

### Bug Fixes 🐛
- **snap:** expose parent directory in device-config plug ([#e81302c](https://github.com/edgexfoundry/device-snmp-go/commits/e81302c))

### Code Refactoring ♻
- **snap:** remove obsolete passthrough usage ([#42a8101](https://github.com/edgexfoundry/device-snmp-go/commits/42a8101))
- **snap:** remove redundant content indentifier ([#3e246a5](https://github.com/edgexfoundry/device-snmp-go/commits/3e246a5))

### Build 👷
- Update to latest SDK w/o ZMQ on windows ([#c9dbc26](https://github.com/edgexfoundry/device-snmp-go/commits/c9dbc26))
    ```
    BREAKING CHANGE:
    ZeroMQ no longer supported on native Windows for EdgeX
    MessageBus
    ```
- **snap:** Source metadata from central repo ([#0413261](https://github.com/edgexfoundry/device-snmp-go/commits/0413261))

### Continuous Integration 🔄
- gomod changes related for Go 1.17 ([#3592ae5](https://github.com/edgexfoundry/device-snmp-go/commits/3592ae5))
- Go 1.17 related changes ([#c9558d6](https://github.com/edgexfoundry/device-snmp-go/commits/c9558d6))

## [v2.1.0] Jakarta - 2021-11-18  (Not Compatible with 1.x releases)

### Bug Fixes 🐛
- update all TOML to use quote and not single-quote ([#0ee2e50](https://github.com/edgexfoundry/device-snmp-go/commits/0ee2e50))

### Documentation 📖
- Update build status badge ([#f9a6796](https://github.com/edgexfoundry/device-snmp-go/commits/f9a6796))

### Build 👷
- Update to use released SDK and go-mods ([#9e6018e](https://github.com/edgexfoundry/device-snmp-go/commits/9e6018e))
- Update alpine base to 3.14 ([#eba3941](https://github.com/edgexfoundry/device-snmp-go/commits/eba3941))

### Continuous Integration 🔄
- Remove need for CI specific Dockerfile ([#7d77ce6](https://github.com/edgexfoundry/device-snmp-go/commits/7d77ce6))

## [v2.0.0] Ireland - 2021-06-30  (Not Compatible with 1.x releases)

### Features ✨
- Enable using MessageBus as the default ([#37e4736](https://github.com/edgexfoundry/device-snmp-go/commits/37e4736))
- Remove Logging configuration ([#9643cb4](https://github.com/edgexfoundry/device-snmp-go/commits/9643cb4))

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

