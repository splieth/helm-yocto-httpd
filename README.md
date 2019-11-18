# helm-yocto-httpd
Deploys [yocto-httpd](https://github.com/felixb/yocto-httpd) as Pod to a K8s cluster. Main purpose anyway is using [terratest](https://github.com/gruntwork-io/terratest) for "unit" and integration testing of helm charts.

## Prerequisites
* [helm](https://helm.sh/)
* [minikube](https://github.com/kubernetes/minikube)
* [go](https://golang.org/)

## Usage
* Spin up a local K8s cluster with minikube and install Tiller: ```make install```
* Run unit test: ```make unittest```
* Run integration test: ```make integrationtest```
* Run all tests: ```make test```
