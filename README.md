# helm-yocto-httpd
Deploys [yocto-httpd](https://github.com/felixb/yocto-httpd) as Pod to a K8s cluster. Main purpose anyway is using [terratest](https://github.com/gruntwork-io/terratest) for "unit" and integration testing of helm charts.

## Prerequisites
* [helm](https://helm.sh/)
* [k3d](https://github.com/k3d-io/k3d)
* [go](https://golang.org/)

## Usage
* Spin up a local K8s (more precisely a [k3s](https://github.com/rancher/k3s)) cluster with k3d: ```make install```
* Run unit test: ```make unittest```
* Run integration test: ```make integrationtest``` (takes about 1.5 minutes)
* Run all tests: ```make test```
