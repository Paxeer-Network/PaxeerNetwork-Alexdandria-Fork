// Copyright PaxLabs Ltd.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/hyperpaxeer-os/blob/main/LICENSE)

package ante_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EvmAnteTestSuite struct {
	suite.Suite
}

func TestEvmAnteTestSuite(t *testing.T) {
	suite.Run(t, &EvmAnteTestSuite{})
}
