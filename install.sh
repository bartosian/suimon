#!/usr/bin/bash

go_version="1.19"
suimon_version="latest"
config_file_name="fullnode.yaml"

sudo apt update && sudo apt upgrade -y && \
sudo apt install wget jq git libclang-dev libpq-dev cmake -y

cd $HOME && \
mkdir -p $HOME/.suimon; \
wget -O $HOME/.suimon/suimon.yaml  https://raw.githubusercontent.com/bartosian/sui_helpers/main/suimon/cmd/checker/config/suimon.template.yaml

if go version | grep -q "$go_version"; then
  echo "Go $go_version is already installed."
else
  echo "Installing Go $go_version..."
  wget "https://dl.google.com/go/go$go_version.linux-amd64.tar.gz"

  if [ -d "/usr/local/" ]; then
    sudo tar -C /usr/local/ -xzf go1.19.linux-amd64.tar.gz && \
    echo "export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin" >> $HOME/.bash_profile && \
    source $HOME/.bash_profile && \
    echo "Go $go_version has been installed successfully."
  else
    if [ -d "/usr/bin/" ]; then
      sudo tar -C /usr/bin/ -xzf go1.19.linux-amd64.tar.gz
      echo "export PATH=$PATH:/usr/bin/go/bin:$HOME/go/bin" >> $HOME/.bash_profile && \
      source $HOME/.bash_profile && \
      echo "Go $go_version has been installed successfully."
    else
      echo "Could not find /usr/local/ or /usr/bin/ directory. Please create one of these directories and run the script again."

      exit 1
    fi
  fi
fi

go install "github.com/bartosian/sui_helpers/suimon@$suimon_version"

result=$(find / -name "$config_file_name" 2>/dev/null)
if [ -z "$result" ]; then
  echo "File not found."
elif [ $(echo "$result" | wc -l) -eq 1 ]; then
  sed -i -e "s%node-config-path:.*%node-config-path: \"$result\"%;" $HOME/.suimon/suimon.yaml
else
  echo "Multiple instances of the $config_file_name found: $result. Please specify path to one of them by using '-nf' flag or 'SUIMON_NODE_CONFIG_PATH' env variable."

  exit 1
fi

echo "Suimon has been installed and configured successfully."
echo
suimon --help