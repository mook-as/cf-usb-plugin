include version.mk

BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)
COMMIT:=$(shell git describe --tags --long | sed -r 's/[0-9.]+-([0-9]+)-(g[a-f0-9]+)/$(VERSION)+\1.\2/')
ARCH:=$(shell go env GOOS).$(shell go env GOARCH)
APP_VERSION=$(COMMIT).$(BRANCH)

PKGSDIRS=$(shell go list -f '{{.Dir}}' ./...)

IMAGE_NAME=helioncf/hcf-cf-plugin-usb
IMAGE_TAG=$(subst +,_,$(APP_VERSION))

print_status = @printf "\033[32;01m==> $(1)\033[0m\n"

.PHONY: all clean format lint vet bindata build test docker docker-write-tag-files docker-push
all: clean format lint vet bindata test build

clean:
	$(call print_status, Cleaning)
	rm -f cf-plugin-usb
	rm -f cf-plugin-usb-*.tgz
	rm -f cf-plugin-usb-*.tar
	rm -f .cf-plugin-usb-docker-*.txt

format:
	$(call print_status, Checking format)
	export GOPATH=$(shell godep path):$(GOPATH) && \
		test 0 -eq `goimports -d -e . | tee /dev/fd/2 | wc -l`

lint:
	$(call print_status, Linting)
	test 0 -eq `echo $(PKGSDIRS) | tr ' ' '\n' | xargs -I '{p}' -n1 golint {p} | tee /dev/fd/2 | wc -l`

vet:
	$(call print_status, Vetting)
	go vet ./...

build:
	$(call print_status, Building)
	export GOPATH=$(shell godep path):$(GOPATH) && \
		go build -ldflags="-X main.version=$(APP_VERSION)"

dist: build
	$(call print_status, Disting)
	tar czf cf-plugin-usb-$(APP_VERSION)-$(ARCH).tgz cf-plugin-usb

tools:
	$(call print_status, Installing Tools)
	go get -u golang.org/x/tools/cmd/vet
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/golang/lint/golint
	go get -u github.com/tools/godep

test:
	$(call print_status, Testing)
	export GOPATH=$(shell godep path):$(GOPATH) &&\
		godep go test -cover ./...

docker:
	$(call print_status, Creating docker image)
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
	docker tag -f $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):$(BRANCH)

# Used by ci
docker-write-tag-files:
	$(call print_status, Writing docker tag files)
	echo $(IMAGE_TAG) > cf-plugin-usb-docker-version-tag.txt
	echo $(BRANCH) > cf-plugin-usb-docker-branch-tag.txt

docker-push: docker
	$(call print_status, Pushing docker image)
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(IMAGE_NAME):$(BRANCH)