#!/bin/bash 

export GO111MODULE=on
export CGO_ENABLED=0
export GOARCH=amd64

#Set image name from skaffold or defaults to the target
export IMAGE=${IMAGE:-$1}

#Set buildcontext from skaffold or default to current dir
export BUILD_CONTEXT=${BUILD_CONTEXT:-.} 

#Set os from cli or default to linux
export GOOS=${2:-linux}

echo "Building $1 with context ${BUILD_CONTEXT}"

if [ "$1" == "email-micro" ]
then
    packr build -o dist/$1 cmd/emailmicroservice/main.go
    docker build -t $IMAGE -f Dockerfile.micro .
else
    echo "Unknown Target"
    exit 1
fi
