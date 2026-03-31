// Copyright PaxLabs Ltd.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/hyperpaxeer-os/blob/main/LICENSE)

package network

import (
	testconstants "github.com/Paxeer-Network/hyperpaxeer-os/testutil/constants"
)

// chainsWHPXHex is an utility map used to retrieve the WHPX contract
// address in hex format from the chain ID.
//
// TODO: refactor to define this in the example chain initialization and pass as function argument
var chainsWHPXHex = map[string]string{
	testconstants.ExampleChainID: testconstants.WHPXContractMainnet,
}

// GetWHPXContractHex returns the hex format of address for the WHPX contract
// given the chainID. If the chainID is not found, it defaults to the mainnet
// address.
func GetWHPXContractHex(chainID string) string {
	address, found := chainsWHPXHex[chainID]

	// default to mainnet address
	if !found {
		address = chainsWHPXHex[testconstants.ExampleChainID]
	}

	return address
}
