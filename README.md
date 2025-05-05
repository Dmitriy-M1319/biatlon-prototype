# Biatlon Prototype

## 1. Build And Run
```bash
make build
```
```bash
./prototype -eventsFile=<event_file> -configFile=<config_json_file>
```

## 2. Structure
* `cmd` - directory with runnable main.go
* `internal/config` - configuration parsing subsystem
* `internal/conveyor` -  subsystem for processing the list of incoming events
* `internal/io` - implementations of interfaces for reading incoming events/writing events and results to the log
* `internal/models` - set of shared structures between subsystems
* `internal/parser` - event parsing subsystem
* `internal/service` - service for calculating individual competitor indicators
