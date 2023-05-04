#!/bin/bash

suimon_version="latest"

# Retrieve the latest stable version of Go from the official website
go_version=$(curl -sSL "https://golang.org/VERSION?m=text")

echo "Installing Suimon..."
echo "======================================"
echo

sudo apt update && sudo apt upgrade -y && \
sudo apt install wget jq git libclang-dev libpq-dev cmake -y

cd $HOME && \
mkdir -p $HOME/.suimon; \
  wget -O $HOME/.suimon/suimon-testnet.yaml https://raw.githubusercontent.com/bartosian/suimon/main/static/suimon.template.yaml

if go version | grep -q "$go_version"; then
  echo "Go $go_version is already installed."
else
  echo "Installing Go $go_version..."
  echo "======================================"
  echo

  wget "https://golang.org/dl/$go_version.linux-amd64.tar.gz"

  if [ -d "/usr/local/" ]; then
    sudo tar -C /usr/local/ -xzf "$go_version.linux-amd64.tar.gz" && \
    echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.bash_profile && \
    source $HOME/.bash_profile && \
    echo "Go $go_version has been installed successfully."
  elif [ -d "/usr/bin/" ]; then
    sudo tar -C /usr/bin/ -xzf "$go_version.linux-amd64.tar.gz" && \
    echo "export PATH=$PATH:/usr/bin/go/bin" >> $HOME/.bash_profile && \
    source $HOME/.bash_profile && \
    echo "Go $go_version has been installed successfully."
  else
    echo "Could not find /usr/local/ or /usr/bin/ directory. Please create one of these directories and run the script again."
    exit 1
  fi
fi

go install "github.com/bartosian/suimon@$suimon_version"

source $HOME/.bash_profile

echo
echo "======================================"
echo "Suimon has been installed and configured successfully."
echo "Before running Suimon, you will need to customize the 'suimon-testnet.yaml' file in the '$HOME/.suimon' directory with the values specific to your environment."
echo "To get started, run 'suimon help'."
