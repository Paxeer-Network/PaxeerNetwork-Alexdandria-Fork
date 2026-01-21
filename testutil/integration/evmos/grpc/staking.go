// Copyright PaxLabs Inc.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/PaxeerNetwork-Alexdandria-Fork/blob/main/LICENSE)
package grpc

import (
	"context"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// GetEvmParams returns the EVM module params.
func (gqh *IntegrationHandler) GetStakingParams() (*stakingtypes.QueryParamsResponse, error) {
	stakingClinet := gqh.network.GetStakingClient()
	return stakingClinet.Params(context.Background(), &stakingtypes.QueryParamsRequest{})
}
