#!/bin/bash

# Retrieve the latest stable version of Go from the official website
go_version=$(curl -sSL "https://golang.org/VERSION?m=text")

echo "Installing Suimon..."
echo "======================================"
echo

# Install Homebrew (if not already installed)
if ! command -v brew &> /dev/null; then
  echo "Homebrew not found. Installing Homebrew..."
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

brew update && brew upgrade && \
brew install wget jq git cmake postgresql && \
brew install llvm

cd $HOME && \
mkdir -p $HOME/.suimon; \
wget -O $HOME/.suimon/suimon-testnet.yaml https://raw.githubusercontent.com/bartosian/suimon/main/static/suimon.template.yaml

if go version | grep -q "$go_version"; then
  echo "Go $go_version is already installed."
else
  echo "Installing Go $go_version..."
  echo "======================================"
  echo

  wget "https://golang.org/dl/$go_version.darwin-amd64.pkg"

  sudo installer -pkg "$go_version.darwin-amd64.pkg" -target /
  echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.zshrc && \
  source $HOME/.zshrc && \
  echo "Go $go_version has been installed successfully."
fi

# Get the latest tag from the GitHub API
LATEST_TAG=$(curl -s https://api.github.com/repos/bartosian/suimon/releases/latest | grep tag_name | cut -d '"' -f 4)

# Download the latest binary release from GitHub
if ! wget "https://github.com/bartosian/suimon/releases/download/$LATEST_TAG/suimon"; then
    echo "Error: Failed to download suimon binary"
    exit 1
fi

# Make the binary executable
chmod +x suimon

# Move the binary to the executable directory
if ! mv suimon /usr/local/bin/; then
    echo "Error: Failed to move suimon binary to /usr/local/bin/"
    exit 1
fi

echo
echo "======================================"
echo "Suimon has been installed and configured successfully."
echo "Before running Suimon, you will need to customize the 'suimon-testnet.yaml' file in the '$HOME/.suimon' directory with the values specific to your environment."
echo "To get started, run 'suimon help'."