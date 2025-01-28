# Collector Distributions

## Installation

### General

#### Environment variables
- `NEW_RELIC_LICENSE_KEY`: New Relic ingest key.
- `NEW_RELIC_MEMORY_LIMIT_MIB`: Maximum amount of memory to be used. 
- `OTEL_EXPORTER_OTLP_ENDPOINT`: New Relic OTLP endpoint to export metrics to, see [official docs](https://docs.newrelic.com/docs/opentelemetry/best-practices/opentelemetry-otlp/)

### Components

The full list of components is available in the respective `manifest.yaml`

### Configuration

The default configuration is `config.yaml` which is embedded in the `Dockerfile` and any OS-specific packaging (if available).
