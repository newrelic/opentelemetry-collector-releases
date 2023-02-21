# Otel collector canaries

This directory contains an Ansible playbook and its dependencies to deploy two
Otel collector containers into a host. The playbook will install Docker and docker-compose into the provided host.

## Usage

Populate the [inventory.yml](./inventory.yml) file with the host/s information.

```bash
$ make PREVIOUS_IMAGE=otel/opentelemetry-collector:0.69.0 CURRENT_IMAGE=otel/opentelemetry-collector:0.70.0
```

A custom inventory file can also be provided:

```bash
$ make INVENTORY_FILE=/a/path PREVIOUS_IMAGE=otel/opentelemetry-collector:0.69.0 CURRENT_IMAGE=otel/opentelemetry-collector:0.70.0
```
