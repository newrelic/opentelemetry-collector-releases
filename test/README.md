## Packaging tests

Run packaging tests:
```shell
# Default values
NR_LICENSE_KEY=nr_license_key make test/packaging

# Limit to linux
NR_LICENSE_KEY=nr_license_key LIMIT=testing_hosts_linux make test/packaging
```
Parameters/Default values:
* `ANSIBLE_INVENTORY_FOLDER`: `PROJECT_ROOT/test/packaging/ansible`
* `ANSIBLE_INVENTORY_FILE`: `inventory.ec2`
* `LIMIT`: `testing_hosts`
* `ANSIBLE_FORKS`: `5`
