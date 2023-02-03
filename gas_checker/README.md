# SUI GAS CHECKER

This tool allows to calculate approximate reference gas price for the next epoch according to the **SuiSystemState**

- [SUI GAS CHECKER](#sui-gas-checker)
    - [Goal](#goal)
    - [Flags](#flags)
    - [Installation Instructions](#installation-instructions)
    - [Usage Examples](#usage-examples)
      - [Getting new calculated gas value every 30 seconds](#getting-new-calculated-gas-value-every-30-seconds)
      - [Getting new calculated gas value with the custom interval(sec)](#getting-new-calculated-gas-value-with-the-custom-intervalsec)
      - [Getting new calculated gas value with the custom RPC](#getting-new-calculated-gas-value-with-the-custom-rpc)
      - [Getting new calculated gas value and updating it on every request if it is different from the last execution](#getting-new-calculated-gas-value-and-updating-it-on-every-request-if-it-is-different-from-the-last-execution)
- [License](#license)

### Goal

- According to current **SuiSystemState** calculate approximate reference gas value for the next epoch

### Flags

| Name      | Default Value                   | Purpose                                                           |
| --------- | ------------------------------- | ----------------------------------------------------------------- |
| rpcURL    | https://fullnode.testnet.sui.io | SUI RPC URL                                                       |
| setGas    | false                           | To determine if validator **NextEpochGasPrice** should be updated |
| frequency | 30                              | Frequency to query and calculate new reference gas price          |

### Installation Instructions

```sh
git clone https://github.com/bartosian/sui_helpers.git
cd sui_helpers/gas_checker

go install

# make sure your PATH variable includes GO install path, which is $HOME/go/bin on most systems. If not set use the following command.
export PATH=$PATH:$HOME/go/bin
```

### Usage Examples

#### Getting new calculated gas value every 30 seconds

```sh
gas_checker
```

#### Getting new calculated gas value with the custom interval(sec)

```sh
gas_checker --frequency 120
```

#### Getting new calculated gas value with the custom RPC

```sh
gas_checker --rpcURL "https://custom.testnet.sui.io"
```

#### Getting new calculated gas value and updating it on every request if it is different from the last execution

```sh
gas_checker --setGas
```

# License

MIT
