#!/bin/bash

apt-get update

if [ $BUILD_FROM_SRC -ne 0 ]
then
  apt-get install -y git curl mercurial

  # deploy bifrost binary
  # && git clone --depth 1 --branch bifrost-v0.0.2 https://github.com/stellar/go.git /go/src/github.com/stellar/go \
  # && git clone --depth 1 --branch horizon-v0.17.3 https://github.com/stellar/go.git /go/src/github.com/stellar/go \

  mkdir -p /go/src/github.com/stellar/ \
    && mv /bifrost/go /go/src/github.com/stellar/ \
    && cd /go/src/github.com/stellar/go \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && dep ensure -v \
    && go install github.com/stellar/go/services/bifrost

    mv /go/bin/bifrost /built-bifrost
 else
  apt-get install -y wget

  wget -O bifrostDL https://github.com/stellar/go/releases/download/bifrost-v0.0.2/bifrost-v0.0.2-linux-amd64
  chmod +x ./bifrostDL

  mv ./bifrostDL /built-bifrost
fi

echo "\nDone installing bifrost...\n"
