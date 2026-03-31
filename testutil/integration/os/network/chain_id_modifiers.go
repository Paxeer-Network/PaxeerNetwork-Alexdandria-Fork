// Copyright PaxLabs Ltd.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/hyperpaxeer-os/blob/main/LICENSE)
//
// This files contains handler for the testing suite that has to be run to
// modify the chain configuration depending on the chainID

package network

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	testconstants "github.com/Paxeer-Network/hyperpaxeer-os/testutil/constants"
	erc20types "github.com/Paxeer-Network/hyperpaxeer-os/x/erc20/types"
)

// updateErc20GenesisStateForChainID modify the default genesis state for the
// bank module of the testing suite depending on the chainID.
func updateBankGenesisStateForChainID(bankGenesisState banktypes.GenesisState) banktypes.GenesisState {
	metadata := generateBankGenesisMetadata()
	bankGenesisState.DenomMetadata = []banktypes.Metadata{metadata}

	return bankGenesisState
}

// generateBankGenesisMetadata generates the metadata
// for the Evm coin depending on the chainID.
func generateBankGenesisMetadata() banktypes.Metadata {
	return banktypes.Metadata{
		Description: "The native EVM, governance and staking token of the HyperPaxeer example chain",
		Base:        "ahpx",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    testconstants.ExampleAttoDenom,
				Exponent: 0,
			},
			{
				Denom:    testconstants.ExampleDisplayDenom,
				Exponent: 18,
			},
		},
		Name:    "HyperPaxeer",
		Symbol:  "PAXEER",
		Display: testconstants.ExampleDisplayDenom,
	}
}

// updateErc20GenesisStateForChainID modify the default genesis state for the
// erc20 module on the testing suite depending on the chainID.
func updateErc20GenesisStateForChainID(chainID string, erc20GenesisState erc20types.GenesisState) erc20types.GenesisState {
	erc20GenesisState.TokenPairs = updateErc20TokenPairs(chainID, erc20GenesisState.TokenPairs)

	return erc20GenesisState
}

// updateErc20TokenPairs modifies the erc20 token pairs to use the correct
// WHPX depending on ChainID
func updateErc20TokenPairs(chainID string, tokenPairs []erc20types.TokenPair) []erc20types.TokenPair {
	testnetAddress := GetWHPXContractHex(chainID)
	coinInfo := testconstants.ExampleChainCoinInfo[chainID]

	mainnetAddress := GetWHPXContractHex(testconstants.ExampleChainID)

	updatedTokenPairs := make([]erc20types.TokenPair, len(tokenPairs))
	for i, tokenPair := range tokenPairs {
		if tokenPair.Erc20Address == mainnetAddress {
			updatedTokenPairs[i] = erc20types.TokenPair{
				Erc20Address:  testnetAddress,
				Denom:         coinInfo.Denom,
				Enabled:       tokenPair.Enabled,
				ContractOwner: tokenPair.ContractOwner,
			}
		} else {
			updatedTokenPairs[i] = tokenPair
		}
	}
	return updatedTokenPairs
}
