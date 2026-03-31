// Copyright PaxLabs Ltd.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/hyperpaxeer-os/blob/main/LICENSE)

package constants

import (
	"fmt"

	evmtypes "github.com/Paxeer-Network/hyperpaxeer-os/x/evm/types"
)

const (
	// DefaultGasPrice is used in testing as the default to use for transactions
	DefaultGasPrice = 20

	// ExampleAttoDenom provides an example denom for use in tests
	ExampleAttoDenom = "ahpx"

	// ExampleMicroDenom provides an example denom for use in tests
	ExampleMicroDenom = "uhpx"

	// ExampleDisplayDenom provides an example display denom for use in tests
	ExampleDisplayDenom = "hpx"

	// ExampleBech32Prefix provides an example Bech32 prefix for use in tests
	ExampleBech32Prefix = "pax"

	// ExampleEIP155ChainID provides an example EIP-155 chain ID for use in tests
	ExampleEIP155ChainID = 125

	// WHPXContractMainnet is the WHPX contract address for mainnet
	WHPXContractMainnet = "0xD4949664cD82660AaE99bEdc034a0deA8A0bd517"
	// WHPXContractTestnet is the WHPX contract address for testnet
	WHPXContractTestnet = "0xcc491f589b45d4a3c679016195b3fb87d7848210"
)

var (
	// ExampleChainIDPrefix provides a chain ID prefix for EIP-155 that can be used in tests
	ExampleChainIDPrefix = fmt.Sprintf("hyperpax_%d", ExampleEIP155ChainID)

	// ExampleChainID provides a chain ID that can be used in tests
	ExampleChainID = ExampleChainIDPrefix + "-1"

	// SixDecimalsChainID provides a chain ID which is being set up with 6 decimals
	SixDecimalsChainID = "hyperpaxsix_126-2"

	// ExampleChainCoinInfo provides the coin info for the example chain
	//
	// It is a map of the chain id and its corresponding EvmCoinInfo
	// that allows initializing the app with different coin info based on the
	// chain id
	ExampleChainCoinInfo = map[string]evmtypes.EvmCoinInfo{
		ExampleChainID: {
			Denom:        ExampleAttoDenom,
			DisplayDenom: ExampleDisplayDenom,
			Decimals:     evmtypes.EighteenDecimals,
		},
		SixDecimalsChainID: {
			Denom:        ExampleMicroDenom,
			DisplayDenom: ExampleDisplayDenom,
			Decimals:     evmtypes.SixDecimals,
		},
	}
)
