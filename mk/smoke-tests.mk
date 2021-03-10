
.PHONY: tests-docker-container smoke-tests start-tests-container stop-tests-container debug-tests-container

TESTS_IMAGE_NAME=smoke-tests:local-dev
TESTS_CONTAINER_ID?=$(shell docker container ls | grep smoke-tests | awk '{print $$1}')
PWD=$(shell pwd)

tests-docker-container: 
	# $(MAKE) build 
	@echo "\n🛠 Building 'kind' 🐳 Docker Image"
	cd smoke-test && cp ../bin/plonk . && docker build . -t $(TESTS_IMAGE_NAME) -f Dockerfile.kind
	cd ..
	@echo "✅ Finished creating 'kind' 🐳 Docker Image\n"

smoke-tests: tests-docker-container 
	@echo "\n🏃‍♂️ Starting 💨 Smoke tests"
	$(MAKE) smoke-test TEST=deploy-test
	@echo "✅ Done running 💨 Smoke tests\n"

smoke-test:
	$(MAKE) start-tests-container TEST=$(TEST)
	@echo "\t🏋️‍♀️ Testing deploy in container: $(TESTS_CONTAINER_ID)"
	smoke-test/scripts/run-smoke-test.sh $(TEST)
	@echo "\t✅ Done testing init\n"
	# $(MAKE) stop-tests-container

start-tests-container:
	@echo "\n🎬 Starting tests container"
	docker run --mount type=bind,source=$(PWD)/smoke-test/$(TEST),target=/app --privileged $(TESTS_IMAGE_NAME) 
	@echo "✅ Finished booting up test container\n"

stop-tests-container:
	@echo "\n🏗 Tearing down the container"
	docker container stop $(TESTS_CONTAINER_ID)
	@echo"✅ Done tearing down the container"

debug-tests-container:
	docker exec --privileged -it `docker container ls | grep smoke-tests | awk '{print $$1}'` /bin/bash
