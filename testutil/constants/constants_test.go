package constants_test

import (
	"testing"

	"github.com/Paxeer-Network/hyperpaxeer-os/testutil/constants"

	"github.com/Paxeer-Network/hyperpaxeer-os/example_chain"
	chainconfig "github.com/Paxeer-Network/hyperpaxeer-os/example_chain/osd/config"
	"github.com/stretchr/testify/require"
)

func TestRequireSameTestDenom(t *testing.T) {
	require.Equal(t,
		constants.ExampleAttoDenom,
		example_chain.ExampleChainDenom,
		"test denoms should be the same across the repo",
	)
}

func TestRequireSameTestBech32Prefix(t *testing.T) {
	require.Equal(t,
		constants.ExampleBech32Prefix,
		chainconfig.Bech32Prefix,
		"bech32 prefixes should be the same across the repo",
	)
}

func TestRequireSameWHPXMainnet(t *testing.T) {
	require.Equal(t,
		constants.WHPXContractMainnet,
		example_chain.WHPXContractMainnet,
		"whpx contract addresses should be the same across the repo",
	)
}
