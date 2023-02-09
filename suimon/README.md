## SUIMON In-Terminal SUI Node Explorer

SUIMON is a terminal explorer for SUI node. The SUIMON explorer displays checkpoints, transactions, uptime, network status, and more information.

## Install SUIMON

The SUIMON installation requires Go. If you don't already have Go installed, see https://golang.org/dl and https://go.dev/doc/install. Download the binary release that is suitable for your system and follow the installation instructions.

### To install the SUIMON binary:

```shell
go install github.com/bartosian/sui_helpers/suimon@latest
```

## Run SUIMON

To launch a SUIMON explorer in your terminal, type:

```shell
suimon -f fullnode.yaml -n testnet
```

## Flags

| Name   | Required | Default | Purpose                                               |
|--------|----------|--------|-------------------------------------------------------|
| ``-f`` | true     |        | path to node config file ``fullnode.yaml``            |
| ``-n`` | false    | devnet | network name. Possible values: ``devnet`` ``testnet`` |

## Print help
```shell
suimon --help
Usage of suimon:
  -f string
    	(required) path to node config file fullnode.yaml
  -n string
    	(optional) network name, possible values: testnet, devnet (default "devnet")
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