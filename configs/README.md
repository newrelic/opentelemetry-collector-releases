# WIP: Configurations for the New Relic Otel collectors

## New Relic Otel collector (nr-otel-collector)

- [nr-otel-collector-agent-linux.yaml](./nr-otel-collector-agent-linux.yaml): Feature parity with the New Relic Infrastructure Agent.

### Usage

#### Package manager

If the collector was installed using a Linux package manager (APT, RPM, etc) some environment variables are predefined in the Systemd service [environment file](../distributions/nr-otel-collector/nr-otel-collector.conf):

```
OTEL_EXPORTER_OTLP_ENDPOINT=otlp.nr-data.net:443
NEW_RELIC_MEMORY_LIMIT_MIB=100
```

The `NEW_RELIC_LICENSE_KEY` environment variable must be set manually, it can be appended to the Systemd service environment file (`/etc/nr-otel-collector/nr-otel-collector.conf`) or direclty to the collector's configuration (`/etc/nr-otel-collector/config.yaml`).

#### Command line

```
OTEL_EXPORTER_OTLP_ENDPOINT=otlp.nr-data.net:4317 NEW_RELIC_MEMORY_LIMIT_MIB=100 NEW_RELIC_LICENSE_KEY=your_license_key nr-otel-collector --config nr-otel-collector-agent-linux.yaml
```
