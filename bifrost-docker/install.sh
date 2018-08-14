#!/bin/bash

apt-get update

apt-get install -y wget

wget -O bifrost https://github.com/stellar/go/releases/download/bifrost-v0.0.2/bifrost-v0.0.2-linux-amd64
chmod +x ./bifrost

mv ./bifrost /bifrost

echo "\nDone installing bifrost...\n"
