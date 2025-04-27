# Container name
CONTAINER_NAME = go-dev

# Build the Docker image
build:
	docker build -t $(CONTAINER_NAME) .

# Run the container with volume mounting
run:
	docker run -d \
		--name $(CONTAINER_NAME) \
		-v $(PWD):/home/app \
		$(CONTAINER_NAME)

# Build and run the container
up: build
	@if docker ps -a | grep -q $(CONTAINER_NAME); then \
		docker rm -f $(CONTAINER_NAME); \
	fi
	@$(MAKE) run

# Enter the container
enter:
	docker exec -it $(CONTAINER_NAME) /bin/sh

# Stop and remove the container
stop:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

# Start an existing container
start:
	docker start $(CONTAINER_NAME)

# Show container status
status:
	docker ps -a | grep $(CONTAINER_NAME)

.PHONY: build run enter stop start status up
