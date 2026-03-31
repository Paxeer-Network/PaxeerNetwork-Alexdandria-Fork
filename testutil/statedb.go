// Copyright PaxLabs Ltd.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/hyperpaxeer-os/blob/main/LICENSE)

package testutil

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	anteinterfaces "github.com/Paxeer-Network/hyperpaxeer-os/ante/interfaces"
	"github.com/Paxeer-Network/hyperpaxeer-os/x/evm/statedb"
)

// NewStateDB returns a new StateDB for testing purposes.
func NewStateDB(ctx sdk.Context, evmKeeper anteinterfaces.EVMKeeper) *statedb.StateDB {
	return statedb.New(ctx, evmKeeper, statedb.NewEmptyTxConfig(common.BytesToHash(ctx.HeaderHash())))
}
