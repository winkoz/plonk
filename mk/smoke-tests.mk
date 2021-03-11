
.PHONY: smoke-tests smoke-test start-tests-container stop-tests-container debug-tests-container tests-container-logs

TESTS_IMAGE_NAME=smoke-tests:local-dev
TESTS_CONTAINER_ID?=$(shell docker container ls | grep smoke-tests | awk '{print $$1}')
PWD=$(shell pwd)

tests-docker-container: 
	# $(MAKE) build 
	@echo "\n🛠 Building 'kind' 🐳 Docker Image"
	cd smoke-test && cp ../bin/plonk . && docker build . -t $(TESTS_IMAGE_NAME)--privileged -f Dockerfile.kind
	cd ..
	@echo "✅ Finished creating 'kind' 🐳 Docker Image\n"

smoke-tests:  
	$(MAKE) start-tests-container
	@echo "\n🏃‍♂️ Starting 💨 Smoke tests"
	$(MAKE) smoke-test TEST=deploy-test
	@echo "✅ Done running 💨 Smoke tests\n"
	$(MAKE) stop-tests-container

smoke-test:
	@echo "\t🏋️‍♀️ Testing \"$(TEST)\" in container: $(TESTS_CONTAINER_ID)"
	smoke-test/scripts/run-smoke-test.sh $(TEST)
	@echo "\t✅ Done testing init\n"

start-tests-container: tests-docker-container
	@echo "\n🎬 Starting tests container"
	rm -rf $(PWD)/smoke-test/sandbox
	mkdir $(PWD)/smoke-test/sandbox
	docker run --mount type=bind,source=$(PWD)/smoke-test/sandbox,target=/app -d --privileged $(TESTS_IMAGE_NAME) 
	@echo "✅ Finished booting up test container\n"

stop-tests-container:
	@echo "\n🏗 Tearing down the container"
	docker container stop $(TESTS_CONTAINER_ID)
	@echo"✅ Done tearing down the container"

debug-tests-container:
	docker exec --privileged -it `docker container ls | grep smoke-tests | awk '{print $$1}'` /bin/bash

tests-container-logs:
	docker logs --follow $(TESTS_CONTAINER_ID)
