<a href="https://opensource.newrelic.com/oss-category/#community-project"><picture><source media="(prefers-color-scheme: dark)" srcset="https://github.com/newrelic/opensource-website/raw/main/src/images/categories/dark/Community_Project.png"><source media="(prefers-color-scheme: light)" srcset="https://github.com/newrelic/opensource-website/raw/main/src/images/categories/Community_Project.png"><img alt="New Relic Open Source community project banner." src="https://github.com/newrelic/opensource-website/raw/main/src/images/categories/Community_Project.png"></picture></a>

# New Relic Distribution of OpenTelemetry (NRDOT) Releases 

This repository assembles various [custom distributions](https://opentelemetry.io/docs/collector/distributions/#custom-distributions) of the [OpenTelemetry Collector](https://opentelemetry.io/docs/collector/) focused on specific use cases and pre-configured to work with NewRelic out-of-the-box.

Generated assets are available in the corresponding Github release page and as docker images published within the [newrelic organization on Docker Hub](https://hub.docker.com/u/newrelic).

Current list of distributions:

- [nrdot-collector-host](./distributions/nrdot-collector-host/): distribution focused on monitoring host metrics and logs
- [nrdot-collector-k8s](./distributions/nrdot-collector-k8s/): distribution focused on monitoring a Kubernetes cluster

## Support

New Relic hosts and moderates an online forum where customers can interact with New Relic employees as well as other customers to get help and share best practices. You can find this project's topic/threads here: [New Relic Community](https://forum.newrelic.com).

## Contribute

We encourage your contributions to improve the New Relic OpenTelemetry collector! Keep in mind that when you submit your pull request, you'll need to sign the CLA via the click-through using CLA-Assistant. You only have to sign the CLA one time per project.

If you have any questions, or to execute our corporate CLA (which is required if your contribution is on behalf of a company), drop us an email at opensource@newrelic.com.

**A note about vulnerabilities**

As noted in our [security policy](../../security/policy), New Relic is committed to the privacy and security of our customers and their data. We believe that providing coordinated disclosure by security researchers and engaging with the security community are important means to achieve our security goals.

If you believe you have found a security vulnerability in this project or any of New Relic's products or websites, we welcome and greatly appreciate you reporting it to New Relic through [HackerOne](https://hackerone.com/newrelic).

If you would like to contribute to this project, review [these guidelines](./CONTRIBUTING.md).

To all contributors, we thank you!  Without your contribution, this project would not be what it is today.

## License
[Collector releases] is licensed under the [Apache 2.0](http://apache.org/licenses/LICENSE-2.0.txt) License.
