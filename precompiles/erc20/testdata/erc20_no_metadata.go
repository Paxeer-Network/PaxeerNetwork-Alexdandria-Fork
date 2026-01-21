// Copyright PaxLabs Inc.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/PaxeerNetwork-Alexdandria-Fork/blob/main/LICENSE)

package testdata

import (
	contractutils "github.com/paxeer/paxeer/v20/contracts/utils"
	evmtypes "github.com/paxeer/paxeer/v20/x/evm/types"
)

func LoadERC20NoMetadataContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ERC20NoMetadata.json")
}
