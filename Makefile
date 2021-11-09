TEST?=$$(go list ./... | grep -v 'vendor')
SWEEP_DIR?=./internal/sweep
SWEEP_ARGS?=
SWEEP_RUN?=all
HOSTNAME=hashicorp.com
NAMESPACE=harness-io
NAME=harness
VERSION=0.2
BINARY=terraform-provider-${NAME}
OS_ARCH=darwin_amd64

default: testacc

# Run acceptance tests
.PHONY: testacc
test:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

# build:
# 	go build -o ${BINARY}
	
# install: build
# 	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
# 	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

# test: 
# 	go test $(TEST) || exit 1                                                   
# 	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4    

sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	go test $(SWEEP_DIR) -v -sweep=all $(SWEEPARGS) -timeout 60m
