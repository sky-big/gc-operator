# Build all binary
#
# Example:
#   make
#   make all
all: build
.PHONY: all

# Build binary
#
# Example:
#   make build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o ./bin/gc-operator ./cmd/controller
.PHONY: build

# Install controller
#
# Example:
#   make install
install:
	deploy/install.sh
.PHONY: install

# UnInstall controller
#
# Example:
#   make uninstall
uninstall:
	deploy/uninstall.sh
.PHONY: uninstall

# Build the docker image
#
# Example:
#   make image
image:
	image/build-image.sh
.PHONY: image

# Push the docker image
#
# Example:
#   make push
push:
	image/push-image.sh
.PHONY: push

# Generate code
#
# Example:
#   make generate
generate:
	hack/codegen/codegen.sh
.PHONY: generate

# Generate vendor
#
# Example:
#   make vendor
vendor:
	go mod vendor
.PHONY: vendor

# Clean package
#
# Example:
#   make clean
clean:
	rm -rf ./bin
