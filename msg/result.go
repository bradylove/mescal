package msg

import (
	"github.com/bradylove/mescal/msg/protocol"
	"io"
	"zombiezen.com/go/capnproto"
)

type SubResult interface {
	getAction() int
	encode(*protocol.Result, *capnp.Segment) error
}

type Result struct {
	SubResult

	Id     string
	Action int
	Status int
}

// NewResult constructs and returns a new Result
func NewResult(id string, status int, sr SubResult) *Result {
	return &Result{sr, id, status, sr.getAction()}
}

// DecodeResult reads from an io.Reader and decodes a result.
func DecodeResult(r io.Reader) (*Result, error) {
	msg, err := capnp.NewDecoder(r).Decode()
	if err != nil {
		return &Result{}, err
	}

	m, err := protocol.ReadRootResult(msg)
	if err != nil {
		return &Result{}, err
	}

	var res Result
	res.Id, err = m.Id()
	if err != nil {
		return &Result{}, err
	}

	switch m.Action() {
	case protocol.Action_handshake:
		if err := res.decodeHandshakeResult(m); err != nil {
			return new(Result), err
		}
	case protocol.Action_get:
		if err := res.decodeGetResult(m); err != nil {
			return new(Result), err
		}
	case protocol.Action_set:
		if err := res.decodeSetResult(m); err != nil {
			return new(Result), err
		}
	default:
		return &Result{}, ErrUnknownAction
	}

	return &res, nil
}

// Encode prepares a result and writes it to an io.Writer
func (rs *Result) Encode(wr io.Writer) error {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return err
	}

	root, err := protocol.NewRootResult(seg)
	if err != nil {
		return err
	}

	err = rs.SubResult.encode(&root, seg)
	if err != nil {
		return err
	}

	root.SetId(rs.Id)
	root.SetStatus(protocol.Status(rs.Status))
	root.SetAction(protocol.Action(rs.SubResult.getAction()))

	err = capnp.NewEncoder(wr).Encode(msg)
	if err != nil {
		return err
	}

	return nil
}

func (rs *Result) decodeHandshakeResult(msg protocol.Result) error {
	sub, err := msg.SubResult().Handshake()
	if err != nil {
		return err
	}

	clientId, err := sub.ClientId()
	if err != nil {
		return err
	}

	rs.SubResult = NewHandshakeResult(clientId)
	rs.Action = rs.SubResult.getAction()

	return nil
}

func (rs *Result) decodeGetResult(msg protocol.Result) error {
	sub, err := msg.SubResult().Get()
	if err != nil {
		return err
	}

	key, err := sub.Key()
	if err != nil {
		return err
	}

	value, err := sub.Value()
	if err != nil {
		return err
	}

	expiry := sub.Expiry()

	rs.SubResult = NewGetResult(key, value, expiry)
	rs.Action = rs.SubResult.getAction()

	return nil
}

func (rs *Result) decodeSetResult(msg protocol.Result) error {
	// sub, err := msg.SubResult().Set()
	// if err != nil {
	//	return err
	// }

	rs.SubResult = NewSetResult()
	rs.Action = rs.SubResult.getAction()

	return nil
}
