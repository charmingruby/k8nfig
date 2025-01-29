DOCKER_IMAGE_NAME = k8nfig
VERSION = 0.1.0

###################
# Docker          #
###################
.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) .

.PHONY: docker-run
docker-run:
	docker run -p $(SERVER_PORT):8080 $(DOCKER_IMAGE_NAME):$(VERSION)
