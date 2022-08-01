HOSTNAME=hashicorp.com
NAMESPACE=Looty
NAME=calculator
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
OS_ARCH=linux_amd64

default: install

build:
	go build -o ${BINARY}

install: build
	rm -rf ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

clean:
	rm -rf ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}

dir:
	ls ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/