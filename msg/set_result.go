package msg

import (
	"github.com/bradylove/mescal/msg/protocol"
	"zombiezen.com/go/capnproto"
)

type SetResult struct{}

func NewSetResult() SetResult {
	return SetResult{}
}

func (gr SetResult) getAction() int {
	return ActionSet
}

func (gr SetResult) encode(root *protocol.Result, s *capnp.Segment) error {
	r, err := protocol.NewResult_SetResult(s)
	if err != nil {
		return err
	}

	root.SubResult().SetSet(r)
	return nil
}
