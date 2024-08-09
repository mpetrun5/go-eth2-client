package capella

import "github.com/mpetrun5/go-eth2-client/spec/capella"

// ExecutionPayloadWithdrawals provides information about withdrawals.
type ExecutionPayloadWithdrawals struct {
	Withdrawals []*capella.Withdrawal `ssz-max:"16"`
}
