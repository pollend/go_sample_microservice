COMMIT_NO := $(shell git rev-parse HEAD 2> /dev/null || true)

BUILD_DIR ?= ./out
$(shell mkdir -p $(BUILD_DIR))

docs:
	docker run --rm -it -e GOPATH=${HOME}/go:/go -v ${HOME}:${HOME} -w $(CURDIR) quay.io/goswagger/swagger generate spec -m -o ./public/docs/swagger.json

flyway:
	@echo "running flyway with: ${ARGS}"
	docker run --rm -v ${PWD}/migration:/flyway/sql flyway/flyway:6.2.1 ${ARGS}

build:
	go build -gcflags='-N -l' -o ./out/app prospect/mobile-physician-api

reload-app:
	docker-compose up -d --no-deps --build app