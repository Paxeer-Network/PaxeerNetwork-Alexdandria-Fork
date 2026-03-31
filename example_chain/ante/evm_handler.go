// Copyright PaxLabs Ltd.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/hyperpaxeer-os/blob/main/LICENSE)

package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	evmante "github.com/Paxeer-Network/hyperpaxeer-os/ante/evm"
)

// newMonoEVMAnteHandler creates the sdk.AnteHandler implementation for the EVM transactions.
func newMonoEVMAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		evmante.NewEVMMonoDecorator(
			options.AccountKeeper,
			options.FeeMarketKeeper,
			options.EvmKeeper,
			options.MaxTxGasWanted,
		),
	)
}
