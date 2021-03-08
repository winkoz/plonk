.PHONY: start-docs-server stop-docs-container debug-docs-container create-site docs-logs

IMAGE=squidfunk/mkdocs-material
WORK_DIR?=
DETACHED=
DOCS_DOCKER?=docker run -it --rm --mount type=bind,source=$(PWD)$(WORK_DIR),target=/src --workdir=/src --privileged $(DETACHED) -p 8000:8000 $(IMAGE)
DOCS_DOCKER_EXEC?=docker exec --workdir=/src --privileged $(DETACHED) $(CONTAINER_ID)
PWD=$(shell pwd)
CONTAINER_ID?=$(shell docker container ls | grep mkdocs | awk '{print $$1}')

create-site: 
	@echo "\n🏗Create docs website infrastructure"
	$(DOCS_DOCKER) mkdocs new plonk-docs
	@echo "✅ Finished creating documentation infrastructure\n"

start-docs-server:
	@echo "\n🎬 Starting mkdocs container"
	$(eval WORK_DIR=/plonk-docs)
	$(eval DETACHED=-d)
	$(DOCS_DOCKER)
	@echo "✅ Finished booting up mkdocs container\n"

stop-docs-container:
	@echo "\n🏗Tearing down the container"
	docker container stop $(CONTAINER_ID)
	@echo"✅ Done tearing down the container"

debug-docs-container:
	docker exec --privileged -it `docker container ls | grep node | awk '{print $$1}'` sh

docs-logs:
	docker logs --follow `docker container ls | grep node | awk '{print $$1}'`
