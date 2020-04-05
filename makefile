SHELL := /bin/bash
MAKEFLAGS += --warn-undefined-variables
.DEFAULT_GOAL := build
.PHONY: *

install-dep:
	@echo "==> Installing dep"
	@go get -u github.com/golang/dep/cmd/dep

install-gometalinter:
	@echo "==> Installing gometalinter"
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

build:
	@echo "==> Building terraform-provider-netbox"
	@go build -mod=vendor -o terraform-provider-netbox .

build-tarball: clean
	@echo "==> Building terraform-provider-netbox"
	@mkdir -p builds
	@if [[ -n $${version} ]]; then export VERSION="_$${version}"; fi ; for GOOS in darwin linux windows; do for GOARCH in 386 amd64; do env GOOS=$${GOOS} GOARCH=$${GOARCH} go build -mod=vendor -o builds/terraform-provider-netbox-$${GOOS}-$${GOARCH}$${VERSION} .; done; done
	@cd builds && for file in *; do echo "$$(sha256sum $${file})" >> ../terraform-provider-netbox_SHA256SUM; done
	@tar czf terraform-provider-netbox.tar.gz builds

localinstall:
	@echo "==> Creating folder terraform.d/plugins/linux_amd64"
	@mkdir -p ~/.terraform.d/plugins/linux_amd64
	@echo "==> Installing provider in this folder"
	@cp terraform-provider-netbox ~/.terraform.d/plugins/linux_amd64

check:
	@echo "==> Checking terraform-provider-netbox"
	@golangci-lint run

clean:
	@echo "==> Cleaning terraform-provider-netbox"
	@rm -f terraform-provider-netbox
	@rm -rf builds
	@rm -f terraform-provider-netbox_SHA256SUM
	@rm -f terraform-provider-netbox.tar.gz
