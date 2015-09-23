package msg

import (
	"github.com/bradylove/mescal/msg/protocol"
	"zombiezen.com/go/capnproto"
)

// GetCommand is get command that can be issued by a client to a server.
type GetCommand struct {
	Key string
}

// NewGetCommand creates a new GetCommand to be used as a SubCommand on a Command.
func NewGetCommand(key string) GetCommand {
	return GetCommand{key}
}

func (c GetCommand) getAction() int {
	return ActionGet
}

func (c GetCommand) encode(root *protocol.Command, s *capnp.Segment) error {
	cmd, err := protocol.NewCommand_GetCommand(s)
	if err != nil {
		return err
	}
	cmd.SetKey(c.Key)

	root.SubCommand().SetGet(cmd)
	return nil
}
