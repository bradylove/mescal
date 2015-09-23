package msg

import (
	"github.com/bradylove/mescal/msg/protocol"
	"zombiezen.com/go/capnproto"
)

type SetCommand struct {
	Key    string
	Value  string
	Expiry int64 // Todo: Should this be Time instead, does capnp suppor that?
}

func NewSetCommand(key, value string, expiry int64) SetCommand {
	return SetCommand{key, value, expiry}
}

func (c SetCommand) getAction() int {
	return ActionSet
}

func (c SetCommand) encode(root *protocol.Command, s *capnp.Segment) error {
	cmd, err := protocol.NewCommand_SetCommand(s)
	if err != nil {
		return err
	}
	cmd.SetKey(c.Key)
	cmd.SetValue(c.Value)
	cmd.SetExpiry(c.Expiry)

	root.SubCommand().SetSet(cmd)
	return nil
}
