package sequencing

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-node/rollup/event"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

type SequencerIface interface {
	event.Deriver
	// NextAction returns when the sequencer needs to do the next change, and iff it should do so.
	NextAction() (t time.Time, ok bool)
	// CheckNextAction is channel to await changes in the value of `nextActionOK`
	CheckNextAction() <-chan struct{}
	Active() bool
	Init(ctx context.Context, active bool) error
	Start(ctx context.Context, head common.Hash) error
	Stop(ctx context.Context) (hash common.Hash, err error)
	SetMaxSafeLag(ctx context.Context, v uint64) error
	OverrideLeader(ctx context.Context) error
	CommitUnsafePayload(*eth.ExecutionPayloadEnvelope) error
	Close()
}
