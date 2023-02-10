## 💧 SUIMON In-Terminal SUI Node Explorer

SUIMON is a terminal explorer for SUI node. The SUIMON explorer displays checkpoints, transactions, uptime, network status, peers, remote RPC and more information.

## Install SUIMON

The SUIMON installation ``requires Go``. If you don't already have Go installed, see https://golang.org/dl and https://go.dev/doc/install. Download the binary release that is suitable for your system and follow the installation instructions.

### To install the SUIMON binary:

```shell
go install github.com/bartosian/sui_helpers/suimon@latest
```

Suimon parses information from the ``fullnode.yaml`` file and its behavior can be configured by using ``suimon.config`` file, ``environment variables`` and ``command line flags``.
Default path where suimon will be looking for any configuration files is ~/.suimon,

### Example suimon.yaml:
```yaml
# update this section if you want to enable/disable tables
monitors-config:
  rpc-table:
    display: true
  node-table:
    display: true
  peers-table:
    display: true

# update this section to add/remove rpc hosts for rpc-table
rpc-config:
  testnet:
    - "https://rpc-office.cosmostation.io/sui-testnet-wave-2"
    - "https://rpc.ankr.com/sui_testnet"
    - "https://sui-api.rpcpool.com"
    - "https://sui-testnet.public.blastapi.io"
    - "https://rpc.ankr.com/sui_testnet"
    - "https://fullnode.testnet.vincagame.com"
    - "https://rpc-office.cosmostation.io/sui-testnet-wave-2"
  devnet:
    - "https://fullnode.devnet.sui.io"

# update this value if you want to enable/disable geolocation
host-lookup-config:
  enable-lookup: true
  geo-db-path: "/Users/admin/.suimon/db"

# update this value to the fullnode.yaml file location
node-config-path: "/Users/admin/.suimon/fullnode.yaml"

# set network to connect to. Possible values: devnet, testmet
network: "testnet"
```

put this file in ``~/.suimon`` directory or use ``-sf`` flag to provide file path to the command or set ``SUIMON_CONFIG_PATH`` environment variable

### Example fullnode.yaml:
```yaml
# Update this value to the location you want Sui to store its database
db-path: "/home/sui/.sui/db"

network-address: "/dns/localhost/tcp/8080/http"
metrics-address: "95.214.54.28:9184"
json-rpc-address: "95.214.54.28:9000"
websocket-address: "95.214.54.28:9001"
enable-event-processing: true

genesis:
  # Update this to the location of where the genesis file is stored
  genesis-file-location: "/home/sui/.sui/genesis.blob"

p2p-config:
  seed-peers:
    - address: "/ip4/65.109.32.171/udp/8084"
    - address: "/ip4/65.108.44.149/udp/8084"
    - address: "/ip4/95.214.54.28/udp/8080"
    - address: "/ip4/136.243.40.38/udp/8080"
    - address: "/ip4/84.46.255.11/udp/8084"
    - address: "/ip4/135.181.6.243/udp/8088"
```

``fullnode.yaml`` file is used by ``SUI`` to configure nodes and validators. ``Suimon`` requires for path to this file being provided. Put this file in ``~/.suimon`` directory or use ``-nf`` flag to provide file path to the command or set ``SUIMON_NODE_CONFIG_PATH`` environment variable

## Run SUIMON

#### Launch SUIMON:

```shell
suimon
```

#### Launch SUIMON and provide suinode.yaml path: 

```shell
suimon -sf suinode.yaml
```

#### Launch SUIMON and provide fullnode.yaml path:

```shell
suimon -nf fullnode.yaml
```

#### Launch SUIMON and provide network name:

```shell
suimon -n testnet
```

## Flags

| Name    | Required | Default               | Purpose                                               |
|---------|----------|-----------------------|-------------------------------------------------------|
| ``-sf`` | false    | path to suinode.yaml  | path to suinode config file ``suinode.yaml``          |
| ``-nf`` | false    | path to fullnode.yaml | path to node config file ``fullnode.yaml``            |
| ``-n``  | false    | devnet                | network name. Possible values: ``devnet`` ``testnet`` |

## Config Files

| Name              | Required | Default Directory | Purpose                                               |
|-------------------|----------|-------------------|-------------------------------------------------------|
| ``suimon.yaml``   | true     | ~/.suimon         | suimon config file                                    |
| ``fullnode.yaml`` | true     | ~/.suimon         | fullnode config file                                  |

## Variables

| Name                        | Required   | Purpose                                               |
|-----------------------------|------------|-------------------------------------------------------|
| ``SUIMON_CONFIG_PATH``      | false      | path to suinode config file ``suinode.yaml``          |
| ``SUIMON_NODE_CONFIG_PATH`` | false      | path to node config file ``fullnode.yaml``            |
| ``SUIMON_NETWORK``          | false      | network name. Possible values: ``devnet`` ``testnet`` |

## Print Help
```shell
suimon --help

Usage of suimon:
  -n string
    	(optional) network name, possible values: testnet, devnet
  -nf string
    	(optional) path to the node config file, can use SUIMON_NODE_CONFIG_PATH variable instead
  -sf string
    	(optional) path to the suimon config file, can use SUIMON_CONFIG_PATH env variable instead
```

## Preview

![Terminal Screenshot](./assets/screenshot_01.png "Screenshot Application")

## Run In Development

To manually run GEX, clone the `github.com/cosmos/gex` repository and then cd into the `gex` directory. Then to run GEX manually, type this command in a terminal window:

```shell
go run -f fullnode.yaml
```

# License

Apache2.0