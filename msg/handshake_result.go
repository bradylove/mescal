package msg

import (
	"github.com/bradylove/mescal/msg/protocol"
	"zombiezen.com/go/capnproto"
)

type HandshakeResult struct {
	ClientId string
}

func NewHandshakeResult(clientId string) HandshakeResult {
	return HandshakeResult{clientId}
}

func (hr HandshakeResult) getAction() int {
	return ActionHandshake
}

func (hr HandshakeResult) encode(root *protocol.Result, s *capnp.Segment) error {
	r, err := protocol.NewResult_HandshakeResult(s)
	if err != nil {
		return err
	}
	r.SetClientId(hr.ClientId)

	root.SubResult().SetHandshake(r)
	return nil
}
