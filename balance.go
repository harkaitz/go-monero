package monero

import (
	"fmt"
)

type Balance struct {
	Balance         XMRAtom `json:"balance"`
	BlocksToUnlock  int64   `json:"blocks_to_unlock"`
	UnlockedBalance XMRAtom `json:"unlocked_balance"`
}

func (m *Monero) GetBalance() (b Balance, err error) {
	err = m.RPCQuery("/json_rpc", "POST", "get_balance", nil, &b)
	return
}

func (b Balance) String() (s string) {
	s += fmt.Sprintf("Locked           : %v\n", b.Balance.ToXMR())
	s += fmt.Sprintf("Unlocked         : %v\n", b.UnlockedBalance.ToXMR())
	s += fmt.Sprintf("Blocks to unlock : %v\n", b.BlocksToUnlock)
	return
}
