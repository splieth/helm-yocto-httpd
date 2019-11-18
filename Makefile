all: unittest integrationtest

fmt:
	gofmt -s -w .

minikube:
	minikube start -p yocto-terratest

tiller:
	helm init --wait

install: minikube tiller

unittest:
	cd test/; go test -run ^TestTemplate

integrationtest:
	cd test/; go test -run ^TestIntegration
