// Copyright PaxLabs Inc.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/PaxeerNetwork-Alexdandria-Fork/blob/main/LICENSE)

package contracts

import (
	contractutils "github.com/paxeer/paxeer/v20/contracts/utils"
	evmtypes "github.com/paxeer/paxeer/v20/x/evm/types"
)

func LoadFlashLoanContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("FlashLoan.json")
}
