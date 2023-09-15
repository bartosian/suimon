#!/bin/bash

# Retrieve the latest stable version of Go from the official website
go_version=$(curl -sSL "https://golang.org/VERSION?m=text")

echo "Installing Suimon..."
echo "======================================"
echo

sudo apt update && sudo apt upgrade -y && \
sudo apt install wget jq git libclang-dev libpq-dev cmake tar -y

cd $HOME && \
mkdir -p $HOME/.suimon; \
  wget -O $HOME/.suimon/suimon-testnet.yaml https://raw.githubusercontent.com/bartosian/suimon/main/static/templates/suimon-testnet.yaml

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
    echo "⚠️ Could not find /usr/local/ or /usr/bin/ directory. Please create one of these directories and run the script again."
    exit 1
  fi
fi

# Define the file name for the release based on the new naming format
RELEASE_FILE="suimon_Linux_x86_64.tar.gz"

# Download the latest binary release from GitHub
if ! wget -O $RELEASE_FILE "https://github.com/bartosian/suimon/releases/download/$LATEST_TAG/$RELEASE_FILE"; then
    echo "⚠️ Error: Failed to download suimon release"
    exit 1
fi

# Extract the binary from the .tar.gz
if ! tar -xzf $RELEASE_FILE; then
    echo "⚠️ Error: Failed to extract suimon binary from $RELEASE_FILE"
    exit 1
fi

# Make the binary executable
chmod +x suimon

# Move the binary to the executable directory
if ! mv suimon /usr/local/bin/; then
    echo "⚠️ Error: Failed to move suimon binary to /usr/local/bin/"
    exit 1
fi

echo
echo "======================================"
echo "🚀 Suimon has been installed and configured successfully. 🎉"
echo "📝 Before running Suimon, you will need to customize the 'suimon-testnet.yaml' file in the '$HOME/.suimon' directory with the values specific to your environment. 🛠️"
echo "👉 To get started, run 'suimon help'. 💡"
