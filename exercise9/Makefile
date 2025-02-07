IMAGE_NAME = myworkout:latest
CONTAINER_NAME = myworkout_container

.PHONY: build
build:
	go build -o bin/myworkout ./cmd/main.go

.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: docker-build
docker-build:
	docker build -t $(IMAGE_NAME) -f docker/Dockerfile .

.PHONY: docker-run
docker-run:
	docker run --name $(CONTAINER_NAME) -p 8080:8080 --rm $(IMAGE_NAME)
