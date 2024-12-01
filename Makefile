DAY ?= $(shell date "+%d")
DAY_DIR ?= ./day/${DAY}

.ONESHELL:

new:
	mkdir -p ${DAY_DIR}
	cd ${DAY_DIR} && \
		go mod init "github.com/tsatam/adventofcode-2024/day/${DAY}" && \
		cd ../../
	go work use ${DAY_DIR}

test:
	go test "${DAY_DIR}/..."

run:
	go run ${DAY_DIR}