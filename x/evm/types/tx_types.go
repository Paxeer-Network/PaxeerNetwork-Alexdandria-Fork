// Copyright PaxLabs Ltd.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/hyperpaxeer-os/blob/main/LICENSE)
package types

import (
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

func GetTxTypeName(txType int) string {
	switch txType {
	case gethtypes.DynamicFeeTxType:
		return "DynamicFeeTxType"
	case gethtypes.LegacyTxType:
		return "LegacyTxType"
	case gethtypes.AccessListTxType:
		return "AccessListTxType"
	default:
		panic("unknown tx type")
	}
}
