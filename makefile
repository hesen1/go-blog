.PHONY: build docker clean tidy lint

all: build

build:
	# @if test -d build; then rm -rf build; fi
	@go build .

docker:
	# @mkdir build
	# @mv blog build/
	# @cp conf/env.ini build/
	# @echo "build 完成"
	docker build --no-cache -t hs/project/blog:$(version) .
	@echo $(version) >> version.txt

clean:
	# @echo "clean"
	# @rm -rf build
	@go clean -i .

tidy:
	@go mod tidy

lint:
	@golint $(path)
