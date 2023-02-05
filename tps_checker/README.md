# SUI TPS CHECKER

This tool allows to calculate TPS on your fullnode and compare it with TPS in the network

- [SUI TPS CHECKER](#sui-tps-checker)
    - [Goal](#goal)
    - [Installation Instructions](#installation-instructions)
    - [Usage Examples](#usage-examples)
    - [Output Examples](#output-examples)
      - [🟢 No Issues](#-no-issues)
      - [🔴 There Are Issues](#-there-are-issues)
- [License](#license)

### Goal

- According to processes **TotalTransactionNumber** in certain time window calculate approximate TPS for your fullnode and in the network

### Installation Instructions

```sh
NETWORK="<testnet | devnet>" # update this variable value to have one of them

wget -O $HOME/check_${NETWORK}_tps.sh https://raw.githubusercontent.com/bartosian/sui_helpers/main/tps_checker/check_${NETWORK}_tps.sh && chmod +x $HOME/check_${NETWORK}_tps.sh

$HOME/check_${NETWORK}_tps.sh
```

### Usage Examples

```sh
$HOME/check_testnet_tps.sh
```

### Output Examples

#### 🟢 No Issues

`number of processed transactions by node is the same as in the network`

```
------------------TESTNET TPS-------------------
Sun 05 Feb 2023 05:06:53 PM CET

IN PROGRESS [100/100] ..........

SUI TPS: 54
NODE TPS: 54
-------------------------------------
```

`when node is not 100% synced with the network, its TPS can be higher (depends on hardware, network and peers)`

```
------------------TESTNET TPS-------------------
Sun 05 Feb 2023 05:06:53 PM CET

IN PROGRESS [100/100] ..........

SUI TPS: 35
NODE TPS: 490
-------------------------------------
```

#### 🔴 There Are Issues

`node's TPS is much lower than in the network (depends on hardware, network and peers)`

```
------------------TESTNET TPS-------------------
Sun 05 Feb 2023 05:06:53 PM CET

IN PROGRESS [100/100] ..........

SUI TPS: 28
NODE TPS: 5 ❗
-------------------------------------
```

`sui remote RPC server doesnt respond`

```
------------------TESTNET TPS-------------------
Sun 05 Feb 2023 05:06:53 PM CET

IN PROGRESS [100/100] ..........

SUI TPS: 0 ❗
NODE TPS: 8
-------------------------------------
```

`node's RPC endpoint does not respond (can be cutom port or node is not running)`

```
------------------TESTNET TPS-------------------
Sun 05 Feb 2023 05:06:53 PM CET

IN PROGRESS [100/100] ..........

SUI TPS: 54
NODE TPS: 0 ❗
-------------------------------------
```

`both metrics equal to 0 (very likely the netwrok is stuck)`

```
------------------TESTNET TPS-------------------
Sun 05 Feb 2023 05:06:53 PM CET

IN PROGRESS [100/100] ..........

SUI TPS: 0 ❗
NODE TPS: 0 ❗
-------------------------------------
```

# License

MIT
