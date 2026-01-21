// Copyright PaxLabs Inc.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/PaxeerNetwork-Alexdandria-Fork/blob/main/LICENSE)
package utils

import (
	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ValidatorConsAddressToHex(valAddress string) common.Address {
	coinbaseAddressBytes := sdk.ConsAddress(valAddress).Bytes()
	return common.BytesToAddress(coinbaseAddressBytes)
}
