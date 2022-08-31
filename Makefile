all: install unittest integrationtest clean

k3d:
	k3d cluster create yocto-terratest

install: k3d

clean:
	k3d cluster delete yocto-terratest

unittest:
	cd test/; go test -run ^TestTemplate

integrationtest:
	cd test/; go test -run ^TestIntegration
