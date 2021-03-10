.DEFAULT_GOAL := build

.PHONY: build

build:
	tinygo build -o ./main.go.wasm -scheduler=none -target=wasi ./main.go

build.docker:
	docker run -it -w /tmp/proxy-wasm-go -v $(shell pwd):/tmp/proxy-wasm-go tinygo/tinygo:0.17.0 \
		tinygo build -o /tmp/proxy-wasm-go/main.go.wasm -scheduler=none -target=wasi \
		/tmp/proxy-wasm-go/main.go

run:
	envoy -c ./envoy.yaml --concurrency 2 --log-format '%v'
