// Copyright PaxLabs Ltd.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/hyperpaxeer-os/blob/main/LICENSE)

package testdata

import (
	contractutils "github.com/Paxeer-Network/hyperpaxeer-os/contracts/utils"
	evmtypes "github.com/Paxeer-Network/hyperpaxeer-os/x/evm/types"
)

func LoadERC20NoMetadataContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ERC20NoMetadata.json")
}
