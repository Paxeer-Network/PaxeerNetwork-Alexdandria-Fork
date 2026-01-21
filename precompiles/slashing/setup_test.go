// Copyright PaxLabs Inc.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/PaxeerNetwork-Alexdandria-Fork/blob/main/LICENSE)

package slashing_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/paxeer/paxeer/v20/precompiles/slashing"
	"github.com/paxeer/paxeer/v20/testutil/integration/evmos/factory"
	"github.com/paxeer/paxeer/v20/testutil/integration/evmos/grpc"
	testkeyring "github.com/paxeer/paxeer/v20/testutil/integration/evmos/keyring"
	"github.com/paxeer/paxeer/v20/testutil/integration/evmos/network"

	"github.com/stretchr/testify/suite"
)

type PrecompileTestSuite struct {
	suite.Suite

	network     *network.UnitTestNetwork
	factory     factory.TxFactory
	grpcHandler grpc.Handler
	keyring     testkeyring.Keyring

	precompile *slashing.Precompile
}

func TestPrecompileTestSuite(t *testing.T) {
	suite.Run(t, new(PrecompileTestSuite))
}

func (s *PrecompileTestSuite) SetupTest() {
	keyring := testkeyring.New(3)
	var err error
	nw := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
		network.WithValidatorOperators([]sdk.AccAddress{
			keyring.GetAccAddr(0),
			keyring.GetAccAddr(1),
			keyring.GetAccAddr(2),
		}),
	)

	grpcHandler := grpc.NewIntegrationHandler(nw)
	txFactory := factory.New(nw, grpcHandler)

	s.network = nw
	s.factory = txFactory
	s.grpcHandler = grpcHandler
	s.keyring = keyring

	if s.precompile, err = slashing.NewPrecompile(
		s.network.App.SlashingKeeper,
		s.network.App.AuthzKeeper,
	); err != nil {
		panic(err)
	}
}
