TEST?=$$(go list ./... | grep -v 'vendor')
default: testacc

build:
	go build -o ${BINARY}

test: 
	go test $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=60s -parallel=4

prep-release:
	./utils/build.sh preprelease
