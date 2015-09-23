package msg

import (
	"github.com/bradylove/mescal/msg/protocol"
	"zombiezen.com/go/capnproto"
)

type GetResult struct {
	Key    string
	Value  string
	Expiry int64
}

func NewGetResult(key, value string, expiry int64) GetResult {
	return GetResult{key, value, expiry}
}

func (gr GetResult) getAction() int {
	return ActionGet
}

func (gr GetResult) encode(root *protocol.Result, s *capnp.Segment) error {
	r, err := protocol.NewResult_GetResult(s)
	if err != nil {
		return err
	}
	r.SetKey(gr.Key)
	r.SetValue(gr.Value)
	r.SetExpiry(gr.Expiry)

	root.SubResult().SetGet(r)
	return nil
}
