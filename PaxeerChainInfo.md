## What is Paxeer?
Paxeer is not a traditional L1 or a single-purpose protocol. It is a capital orchestration network designed to implement a capital-to-user pipeline: a system where protocol-level balance sheet, risk, and execution are composed into a single substrate that continuously allocates on-chain resources to users, agents, and strategies.

## Capital-to-user model
The network that funds you
Paxeer generalizes the “prop firm” construct from isolated trading accounts into a network-level funding plane. Instead of capital being siloed inside a single broker or account, a protocol-governed pool of resources is continuously allocated across builders, traders, liquidity providers, and autonomous agents operating on-chain.

In practice, Paxeer composes ChainFlow V2 as a capital routing and risk assessment engine, Paxeer itself as the deterministic settlement layer, and OpenNet as an external execution environment into a single pipeline: capital → risk model → user wallet → on-chain action.

From prop firm to protocol

ChainFlow V2 is no longer a discrete “prop trading firm”. It is a protocol surface inside Paxeer where the network provides the capital and users interact through funded smart wallets, without auditions, upfront fees, or opaque allocation rules.

## Account surface
Protocol-funded smart wallets
Instead of provisioning monolithic “prop accounts”, Paxeer issues ChainFlow smart wallets to users. These are programmatic accounts whose initial and ongoing funding is derived from the network's capital pool, not a single institution. Each wallet is treated as a first-class capital conduit, able to: deploy applications, trade on any on-chain protocol, provide liquidity, or hold portfolio allocations.

All interactions are settled transparently on Paxeer and OpenNet, with the smart wallet acting as a programmable perimeter around user behavior. This creates a clear separation between who controls actions (the user or agent) and who supplies capital (the protocol and its funding pools).

## View OpenNet & Paxeer activity on explorer →
Community balance sheet
Network-level funding pools
Capital originates from community funding pools denominated in ETH and OP, economically backed by approximately 1.5B USD worth of collateralized, staked PAX coins. Rather than assigning static balances, the protocol continuously routes this capital across users, strategies, and risk buckets, ensuring that size is allocated where risk-adjusted utility is highest.

At equilibrium, Paxeer behaves like a multi-surface capital router: liquidity is treated as a shared resource, and the network decides how much capital any given wallet is entitled to deploy at a given time, based on observed behavior and aggregate system constraints.

Multi-surface
Trade, build, invest — one stack
The capital-to-user pipeline is surface-agnostic. A funded wallet can be used to trade spot or derivatives, deploy smart contracts, provide AMM liquidity, underwrite credit, or compose new protocols. From the network's perspective, all of these are just state transitions with different risk signatures.

This collapses the distinction between “prop trading”, “builder grants”, and “investment capital” into a single programmable capital rail, mediated by the same risk engine and settlement layer.

## No auditions
No fees, targets, or challenges
Traditional prop firms gate access via evaluations, drawdown rules, and profit targets. Paxeer removes this layer entirely. There are no audition fees, no profit-challenge ladders, and no externalized evaluation products. Users connect and transact as they normally would; evaluation is implicit in how the risk engine updates capital limits and payout schedules.

Operationally, this feels like using a normal wallet, except that the underlying balance sheet is owned by the network and continuously sized by protocol rules rather than by the individual user.

## Capital semantics
Network-level capital model
Paxeer treats capital as a first-class network primitive. Instead of balances being an emergent artifact of user deposits, funded capacity is explicit, revocable, and stateful across sessions. This enables higher-order semantics like per-user risk envelopes, path-dependent limits, and programmatic capital escalation based on track record.

## Algorithmic payouts
LLM-powered risk & payout engine
Underneath the user experience is a proprietary, network-wide risk and payout engine. It continuously ingests more than 500 on-chain and behavioral data points per funded wallet: position profiles, PnL paths, latency patterns, protocol mix, liquidity usage, temporal clustering of actions, and more.

Large Language Models (LLMs) and domain-specific risk models operate as an ensemble: LLMs interpret high-dimensional behavioral traces and strategy descriptors, while quantitative models enforce hard limits, drawdown envelopes, and solvency constraints. The result is a dynamic entitlement function that determines when and how much a user can withdraw, and how much additional capital they can safely deploy.

Conceptually, the engine turns “normal on-chain behavior” into executable credit: your usage pattern becomes a real-time signal that the network converts into funded capacity and automated, transparent payouts.

## High-level schematic

Community pools lock ETH/OP collateral against staked PAX.
Risk engine computes per-wallet envelopes and funded capacity.
ChainFlow smart wallets execute actions on Paxeer + OpenNet.
Engine re-scores behavior and updates limits and payout queues.
Payouts are settled on-chain back to users or contributors.
In one sentence
Paxeer is a capital-to-user execution network: a chain where protocol-governed balance sheet, LLM-driven risk models, and funded smart wallets are fused into a single system that allocates and settles capital for anyone building, trading, or experimenting on-chain.