<!--
parent:
  order: false
-->

<div align="center">
  <h1> HyperPaxeer Network </h1>
</div>

<div align="center">
  <a href="https://github.com/paxeer/hyperpaxeer/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/paxeer/hyperpaxeer.svg" />
  </a>
  <a href="https://github.com/paxeer/hyperpaxeer/blob/main/LICENSE">
    <img alt="License" src="https://img.shields.io/github/license/paxeer/hyperpaxeer.svg" />
  </a>
</div>

## About

HyperPaxeer Network is a capital orchestration network built as part of the Alexandria Fork upgrade, 
migrating validators from the original Paxeer Network (geth PoS, chain ID 229) to this new 
high-throughput Proof-of-Stake EVM blockchain (chain ID 125).

It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/)
which runs on top of the [CometBFT](https://github.com/cometbft/cometbft) consensus engine,
providing full Ethereum compatibility and interoperability.

### Network Information

| Parameter | Value |
|-----------|-------|
| Chain ID (Cosmos) | `hyperpax_125-1` |
| EVM Chain ID | `125` |
| Token Symbol | `PAX` |
| Base Denomination | `ahpx` |
| Display Denomination | `hpx` |
| Bech32 Prefix | `pax` |
| RPC Endpoint | `https://mainnet-beta.rpc.hyperpaxeer.com/rpc` |
| Block Explorer | `https://paxscan.paxeer.app` |

## Quick Start

### For New Validators

1. Clone this repository
2. Install the binary:
```bash
make install
```

3. Initialize your node:
```bash
hyperpaxd init <your-moniker> --chain-id hyperpax_125-1
```

4. Download the genesis file from the network and place it in `~/.hyperpaxd/config/genesis.json`

5. Configure your node with the appropriate peers and seeds

6. Start your node:
```bash
hyperpaxd start
```

## Installation

For prerequisites and detailed build instructions, ensure you have Go 1.22+ installed.
Once the dependencies are installed, run:

```bash
make install
```

This will install the `hyperpaxd` binary to your `$GOPATH/bin`.

## Configuration

See the following configuration files for network setup:
- `config.env` - Network configuration parameters
- `shared/network_info.json` - Network validator information

## Documentation

For more information about Paxeer and HyperPaxeer Network, see [PaxeerChainInfo.md](./PaxeerChainInfo.md).

## Community

- Website: [hyperpaxeer.com](https://hyperpaxeer.com)
- Block Explorer: [paxscan.paxeer.app](https://paxscan.paxeer.app)

## Licensing

This software is based on Paxeer Network and is subject to the [Paxeer Network Non-Commercial License 1.0 (ENCL-1.0)](./LICENSE).
