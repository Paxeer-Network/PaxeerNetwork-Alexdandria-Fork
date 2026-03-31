// Copyright PaxLabs Ltd.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/hyperpaxeer-os/blob/main/LICENSE)

package example_chain

import erc20types "github.com/Paxeer-Network/hyperpaxeer-os/x/erc20/types"

// WHPXContractMainnet is the WHPX contract address for mainnet
const WHPXContractMainnet = "0xD4949664cD82660AaE99bEdc034a0deA8A0bd517"

// ExampleTokenPairs creates a slice of token pairs, that contains a pair for the native denom of the example chain
// implementation.
var ExampleTokenPairs = []erc20types.TokenPair{
	{
		Erc20Address:  WHPXContractMainnet,
		Denom:         ExampleChainDenom,
		Enabled:       true,
		ContractOwner: erc20types.OWNER_MODULE,
	},
}
