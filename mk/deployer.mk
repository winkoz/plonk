.PHONY: build-deployer publish-deployer docker-login
DOCKER_REGISTRY?=registry.winkoz.com
DOCKER_USER?=winkoz

build-deployer:
	docker build -f Dockerfile.deployer -t registry.winkoz.com/plonk_deployer:latest .

docker-login:
	docker login registry.winkoz.com -u $(DOCKER_USER) -p $(DOCKER_PASS)

publish-deployer: build-deployer docker-login
	docker push $(DOCKER_REGISTRY)/plonk_deployer:latest