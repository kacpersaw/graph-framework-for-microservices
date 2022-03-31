#
# App Info
#
APP_NAME ?= api-gw
NAMESPACE ?= default
DATAMODEL_DIR ?= "./nexus"
CLUSTER ?= ""
DATAMODEL ?= ""
DATAMODEL_GROUP ?= ""
NEXUS_BIN ?= $(shell which nexus)

#
# Image Info
#
### adding this to test app init..
CI_COMMIT ?= $(shell git rev-parse --verify --short=8 HEAD 2> /dev/null || echo "00000000")
IMAGE_TAG ?= ${APP_NAME}-${CI_COMMIT}
GIT_HEAD  ?= $(shell git rev-parse --verify HEAD 2> /dev/null || echo "0000000000000000")
IMAGE_REGISTRY ?= 284299419820.dkr.ecr.us-west-2.amazonaws.com/nexus/api-gw

#
# Platform
#
NEXUS_CLI_TAG ?= latest

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php


.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Platform
.PHONY: nexus-cli
nexus-cli: ## Install Nexus CLI
	curl https://storage.googleapis.com/nexus-cli-downloads/nexus-$(uname | tr '[:upper:]' '[:lower:]')_amd64 --output nexus
	chmod 755 nexus
	mv nexus /usr/local/bin/nexus

## to silence all not needed.
%:
	@true

replace:
	if [[ -n ${DATAMODEL} ]]; then\
			test -s ${DATAMODEL_DIR}/${DATAMODEL} || { echo "Please create datamodel ${DATAMODEL} for go mod replace"; exit 1; } ;\
			go mod edit -replace $(DATAMODEL)=${DATAMODEL_DIR}/${DATAMODEL} ;\
	fi

.SILENT:
.PHONY: datamodel_init
datamodel_init: ## Initialize datamodel
	if [ -z $(NEXUS_BIN) ]; then \
		echo "Please install nexus binary"; \
		exit 1; \
	fi 
	if [[ -n ${DATAMODEL} ]]; then \
		if [[ -n ${DATAMODEL_GROUP} ]]; then \
			$(NEXUS_BIN) datamodel init --name ${DATAMODEL} --group ${DATAMODEL_GROUP};\
		else \
			$(NEXUS_BIN) datamodel init --name ${DATAMODEL} ;\
		fi \
	else \
		$(NEXUS_BIN) datamodel init ;\
	fi
	$(MAKE) replace

##@ Dev

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: lint
lint: ## lint checks using the make targets
	$(MAKE) fmt
	$(MAKE) vet

.PHONY: build
build: lint ## Build manager binary.
	docker build --build-arg APP_NAME=${APP_NAME} \
				--build-arg GIT_HEAD=${GIT_HEAD} \
				--build-arg GIT_TAG=${CI_COMMIT} \
				--build-arg CICD_TOKEN=${CICD_TOKEN} \
				-t ${IMAGE_REGISTRY}:${IMAGE_TAG} .

##@ Test

.PHONY: test
test:
	go test ./...

.PHONY: integration_test
integration_test:
	echo "Add your integration_tests for your app here!!!!..." ;
	echo "Sample Workflow can be.";
	echo "1. Create Environment";
	echo "2. Start built application in the Environment";
	echo "3. Start integration_tests with go test / gingko framework";

.PHONY: integration_test_cleanup
integration_test_cleanup:
	echo "Add your cleanup steps here!!!!....";
	echo "Possible steps you can do.";
	echo "1. Get logs of integration test as artifacts"
	echo "2. Get logs of components in clusters as artifacts for debugging"

.PHONY: teardown_environment
teardown_environment:
	echo "Add cluster cleanup step after integration_tests pass/fail here..";
	echo "Clear clusters created";

##@ Coverage
.PHONY: coverage
coverage:
    go test -json -coverprofile=coverage.out ./... | tee report.json ;\
	sonar-scanner ;

##@ Publish

.PHONY: publish
publish: build ## Push docker image with the manager.
	docker push ${IMAGE_REGISTRY}:${IMAGE_TAG}

##@ Security Scan

.PHONY: image_scan
image_scan:
	flash docker scan image ${IMAGE_REGISTRY}:${IMAGE_TAG}

##@ Deploy

.PHONY: deploy
deploy: kustomize
	if [[ $(CLUSTER) == "kind" ]]; then \
		kind load docker-image ${IMAGE_REGISTRY}:${IMAGE_TAG} ;\
	fi
	cd config/deployment/ && $(KUSTOMIZE) edit set image ${APP_NAME}=${IMAGE_REGISTRY}:${IMAGE_TAG} && $(KUSTOMIZE) build . | kubectl apply -f - -n ${NAMESPACE};


.PHONY: undeploy
undeploy: kustomize
	cd config/deployment/ && $(KUSTOMIZE) build . | kubectl delete -f - -n ${NAMESPACE};


#check how to use kustomize for now using sed to replace deployment..
KUSTOMIZE = $(shell pwd)/bin/kustomize
.PHONY: kustomize
kustomize: ## Download kustomize locally if necessary.
	$(MAKE) install-kustomize

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
install-kustomize:
	test -s $(PROJECT_DIR)/bin/kustomize || { mkdir -p $(PROJECT_DIR)/bin; cd $(PROJECT_DIR)/bin; curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash ; };

