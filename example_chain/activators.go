// Copyright PaxLabs Ltd.(Paxeer Network)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/Paxeer-Network/hyperpaxeer-os/blob/main/LICENSE)
package example_chain

import (
	"github.com/Paxeer-Network/hyperpaxeer-os/example_chain/eips"
	"github.com/Paxeer-Network/hyperpaxeer-os/x/evm/core/vm"
)

// paxeerActivators defines a map of opcode modifiers associated
// with a key defining the corresponding EIP.
var paxeerActivators = map[string]func(*vm.JumpTable){
	"paxeer_0": eips.Enable0000,
	"paxeer_1": eips.Enable0001,
	"paxeer_2": eips.Enable0002,
}
