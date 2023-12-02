IMAGE_NAME = linkster
DOCKER ?= 1

.PHONY: build run

build:
ifeq ($(DOCKER), 1)
	docker build -t $(IMAGE_NAME) .
else
	@go build -o linkster
endif

run: build
ifeq ($(DOCKER), 1)
	docker run --network host $(IMAGE_NAME)
else
	./linkster
endif

test:
	go test -v ./...

clean:
ifeq ($(DOCKER), 1)
	docker rmi -f $(IMAGE_NAME)
else
	rm linkster
endif