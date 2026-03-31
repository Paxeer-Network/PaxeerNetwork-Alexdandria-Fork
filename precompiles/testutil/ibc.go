// Copyright PaxLabs Ltd.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/hyperpaxeer-os/blob/main/LICENSE)

package testutil

import (
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
)

var (
	UosmoDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: "uosmo",
	}
	UosmoIbcdenom = UosmoDenomtrace.IBCDenom()

	UatomDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-1",
		BaseDenom: "uatom",
	}
	UatomIbcdenom = UatomDenomtrace.IBCDenom()

	UpaxeerDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: "ahpx",
	}
	UpaxeerIbcdenom = UpaxeerDenomtrace.IBCDenom()

	UatomOsmoDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-0/transfer/channel-1",
		BaseDenom: "uatom",
	}
	UatomOsmoIbcdenom = UatomOsmoDenomtrace.IBCDenom()

	ApaxeerDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: "ahpx",
	}
	ApaxeerIbcdenom = ApaxeerDenomtrace.IBCDenom()
)
