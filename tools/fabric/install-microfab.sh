#!/bin/bash

export DIR="$(dirname -- "$(readlink -f "${BASH_SOURCE}")")"

cd $DIR

curl -sSL https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh  | bash -s -- binary

# ${DIR}/install-fabric.sh

curl -sSL https://github.com/hyperledger-labs/microfab/releases/download/v0.0.18/microfab-linux-amd64 -o bin/microfab

echo "export PATH=$PATH:${pwd}/bin" >> ~/.bashrc
echo "export PATH=$PATH:${pwd}/bin" >> ~/.zshrc
echo "export FABRIC_CFG_PATH=$(pwd)/config" >> ~/.bashrc
echo "export FABRIC_CFG_PATH=$(pwd)/config" >> ~/.zshrc
