# HyperPaxeer Network - Validator Setup Guide

This guide walks you through setting up a validator node on the HyperPaxeer Network (Chain ID: 125).

## Prerequisites

- **OS**: Ubuntu 20.04+ or similar Linux distribution
- **Hardware**:
  - CPU: 4+ cores
  - RAM: 16GB minimum (32GB recommended)
  - Storage: 500GB+ NVMe SSD
  - Network: 100Mbps+ stable connection
- **Software**:
  - Go 1.22+
  - Git
  - Make
  - jq

## Network Information

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

## Step 1: Install Dependencies

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install required packages
sudo apt install -y build-essential git curl jq lz4

# Install Go 1.22+
wget https://go.dev/dl/go1.22.8.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.8.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verify Go installation
go version
```

## Step 2: Build and Install hyperpaxd

```bash
# Clone the repository
git clone https://github.com/paxeer/hyperpaxeer.git
cd hyperpaxeer

# Build and install
make install

# Verify installation
hyperpaxd version
```

## Step 3: Initialize Your Node

```bash
# Set your moniker (validator name)
MONIKER="your-validator-name"

# Initialize the node
hyperpaxd init $MONIKER --chain-id hyperpax_125-1

# This creates the default configuration in ~/.hyperpaxd/
```

## Step 4: Configure Genesis

Download the genesis file from the network:

```bash
# Download genesis.json
curl -L https://mainnet-beta.rpc.hyperpaxeer.com/genesis -o ~/.hyperpaxd/config/genesis.json

# Verify genesis file
hyperpaxd validate-genesis
```

## Step 5: Configure Peers and Seeds

Edit `~/.hyperpaxd/config/config.toml`:

```bash
# Get persistent peers from existing validators
PEERS="<node_id>@<ip>:26656,<node_id>@<ip>:26656"

# Update config.toml
sed -i "s/persistent_peers = \"\"/persistent_peers = \"$PEERS\"/" ~/.hyperpaxd/config/config.toml
```

## Step 6: Configure App Settings

Edit `~/.hyperpaxd/config/app.toml`:

```bash
# Set minimum gas price
sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.0001ahpx"/' ~/.hyperpaxd/config/app.toml

# Enable API (optional)
sed -i 's/enable = false/enable = true/' ~/.hyperpaxd/config/app.toml

# Enable JSON-RPC for EVM compatibility
sed -i 's/address = "127.0.0.1:8545"/address = "0.0.0.0:8545"/' ~/.hyperpaxd/config/app.toml
```

## Step 7: Create or Import Wallet

```bash
# Create a new wallet
hyperpaxd keys add validator --keyring-backend file

# OR import existing wallet from mnemonic
hyperpaxd keys add validator --recover --keyring-backend file

# Save the mnemonic phrase securely!
```

## Step 8: Start the Node

### Option A: Direct Start (for testing)

```bash
hyperpaxd start \
    --minimum-gas-prices=0.0001ahpx \
    --json-rpc.api eth,txpool,personal,net,debug,web3
```

### Option B: Systemd Service (recommended for production)

Create `/etc/systemd/system/hyperpaxd.service`:

```ini
[Unit]
Description=HyperPaxeer Node
After=network-online.target

[Service]
User=$USER
ExecStart=$(which hyperpaxd) start --minimum-gas-prices=0.0001ahpx --json-rpc.api eth,txpool,personal,net,debug,web3
Restart=on-failure
RestartSec=3
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
```

Enable and start:

```bash
sudo systemctl daemon-reload
sudo systemctl enable hyperpaxd
sudo systemctl start hyperpaxd

# Check logs
sudo journalctl -u hyperpaxd -f
```

## Step 9: Wait for Sync

Your node needs to sync with the network before creating a validator:

```bash
# Check sync status
hyperpaxd status 2>&1 | jq .SyncInfo.catching_up

# When catching_up is false, you're synced
```

## Step 10: Create Validator

Once synced and you have sufficient HPX tokens:

```bash
hyperpaxd tx staking create-validator \
    --amount=1000000000000000000000ahpx \
    --pubkey=$(hyperpaxd tendermint show-validator) \
    --moniker="$MONIKER" \
    --chain-id=hyperpax_125-1 \
    --commission-rate="0.05" \
    --commission-max-rate="0.20" \
    --commission-max-change-rate="0.01" \
    --min-self-delegation="1000000000000000000" \
    --gas="auto" \
    --gas-adjustment=1.5 \
    --gas-prices="0.0001ahpx" \
    --from=validator \
    --keyring-backend=file
```

## Useful Commands

### Check Validator Status
```bash
hyperpaxd query staking validator $(hyperpaxd keys show validator --bech val -a --keyring-backend file)
```

### Delegate More Tokens
```bash
hyperpaxd tx staking delegate $(hyperpaxd keys show validator --bech val -a --keyring-backend file) <amount>ahpx \
    --from=validator --chain-id=hyperpax_125-1 --gas=auto --keyring-backend=file
```

### Unjail Validator
```bash
hyperpaxd tx slashing unjail --from=validator --chain-id=hyperpax_125-1 --gas=auto --keyring-backend=file
```

### Check Balance
```bash
hyperpaxd query bank balances $(hyperpaxd keys show validator -a --keyring-backend file)
```

## EVM/MetaMask Configuration

To connect MetaMask or other EVM wallets:

| Setting | Value |
|---------|-------|
| Network Name | HyperPaxeer Network |
| RPC URL | `https://mainnet-beta.rpc.hyperpaxeer.com/rpc` |
| Chain ID | `125` |
| Currency Symbol | `PAX` |
| Block Explorer | `https://paxscan.paxeer.app` |

## Troubleshooting

### Node Not Syncing
- Check peers: `hyperpaxd net info`
- Verify genesis hash matches network
- Check firewall allows port 26656

### Out of Memory
- Increase swap or RAM
- Enable pruning in `app.toml`

### Validator Jailed
- Check slashing reason: `hyperpaxd query slashing signing-info $(hyperpaxd tendermint show-validator)`
- Wait for jail period, then unjail

## Support

- Block Explorer: [paxscan.paxeer.app](https://paxscan.paxeer.app)
- Documentation: See [PaxeerChainInfo.md](./PaxeerChainInfo.md)
