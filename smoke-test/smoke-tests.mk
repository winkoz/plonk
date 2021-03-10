
.PHONY: docker-container smoke-tests start-container stop-container smoke-test debug-container

IMAGE_NAME=smoke-tests:local-dev
CONTAINER_ID?=$(shell docker container ls | grep smoke-tests | awk '{print $$1}')
PWD=$(shell pwd)

docker-container: 
	# $(MAKE) build 
	@echo "\n🛠 Building 'kind' 🐳 Docker Image"
	cd smoke-test && cp ../bin/plonk . && docker build . -t $(IMAGE_NAME) -f Dockerfile.kind
	cd ..
	@echo "✅ Finished creating 'kind' 🐳 Docker Image\n"

smoke-tests: docker-container 
	@echo "\n🏃‍♂️ Starting 💨 Smoke tests"
	$(MAKE) smoke-test TEST=deploy-test
	@echo "✅ Done running 💨 Smoke tests\n"

smoke-test:
	$(MAKE) start-container TEST=$(TEST)
	@echo "\t🏋️‍♀️ Testing deploy in container: $(CONTAINER_ID)"
	smoke-test/scripts/run-smoke-test.sh $(TEST)
	@echo "\t✅ Done testing init\n"
	# $(MAKE) stop-container

start-container:
	@echo "\n🎬 Starting tests container"
	docker run --mount type=bind,source=$(PWD)/smoke-test/$(TEST),target=/app --privileged $(IMAGE_NAME) 
	@echo "✅ Finished booting up test container\n"

stop-container:
	@echo "\n🏗 Tearing down the container"
	docker container stop $(CONTAINER_ID)
	@echo"✅ Done tearing down the container"

debug-container:
	docker exec --privileged -it `docker container ls | grep smoke-tests | awk '{print $$1}'` /bin/bash
