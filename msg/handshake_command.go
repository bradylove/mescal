package msg

import (
	"github.com/bradylove/mescal/msg/protocol"
	"zombiezen.com/go/capnproto"
)

// HandshakeCommand is the first command that should be issued to a server
// imediately after opening a new connection.
type HandshakeCommand struct {
	UserAgent string
}

// NewHandshakeCommand creates a new HandshakeCommand to be used as a SubCommand
// on a Command.
func NewHandshakeCommand(agent string) HandshakeCommand {
	return HandshakeCommand{agent}
}

func (c HandshakeCommand) getAction() int {
	return ActionHandshake
}

func (c HandshakeCommand) encode(root *protocol.Command, s *capnp.Segment) error {
	cmd, err := protocol.NewCommand_HandshakeCommand(s)
	if err != nil {
		return err
	}
	cmd.SetUserAgent(c.UserAgent)

	root.SubCommand().SetHandshake(cmd)
	return nil
}
