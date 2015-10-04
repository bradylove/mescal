package msg

import (
	"github.com/bradylove/mescal/msg/protocol"
	"io"
	"zombiezen.com/go/capnproto"
)

type SubCommand interface {
	getAction() int
	encode(*protocol.Command, *capnp.Segment) error
}

// Command is the root element for issuing a command from a client to a server.
type Command struct {
	SubCommand

	Id     string
	Action int
}

// NewCommand constructs and returns a new Command.
func NewCommand(id string, sb SubCommand) *Command {
	return &Command{sb, id, sb.getAction()}
}

// DecodeCommand reads from an io.Reader and decodes the command.
func DecodeCommand(r io.Reader) (*Command, error) {
	msg, err := capnp.NewDecoder(r).Decode()
	if err != nil {
		return &Command{}, err
	}

	m, err := protocol.ReadRootCommand(msg)
	if err != nil {
		return &Command{}, err
	}

	var cmd Command
	cmd.Id, err = m.Id()
	if err != nil {
		return &Command{}, err
	}

	switch m.Action() {
	case protocol.Action_handshake:
		if err := cmd.decodeHandshakeCommand(m); err != nil {
			return &Command{}, err
		}
	case protocol.Action_get:
		if err := cmd.decodeGetCommand(m); err != nil {
			return &Command{}, err
		}
	case protocol.Action_set:
		if err := cmd.decodeSetCommand(m); err != nil {
			return &Command{}, err
		}
	default:
		return &Command{}, ErrUnknownAction
	}

	return &cmd, nil
}

// Encode prepares a command to be issued over the wire.
func (c *Command) Encode(wr io.Writer) error {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return err
	}

	root, err := protocol.NewRootCommand(seg)
	if err != nil {
		return err
	}

	err = c.SubCommand.encode(&root, seg)
	if err != nil {
		return err
	}

	root.SetId(c.Id)
	root.SetAction(protocol.Action(c.SubCommand.getAction()))

	err = capnp.NewEncoder(wr).Encode(msg)
	if err != nil {
		return err
	}

	return nil
}

func (c *Command) decodeSetCommand(msg protocol.Command) error {
	subCmd, err := msg.SubCommand().Set()
	if err != nil {
		return err
	}

	key, err := subCmd.Key()
	if err != nil {
		return err
	}

	value, err := subCmd.Value()
	if err != nil {
		return err
	}

	expiry := subCmd.Expiry()
	if err != nil {
		return err
	}

	c.SubCommand = NewSetCommand(key, value, expiry)
	c.Action = c.SubCommand.getAction()
	return nil
}

func (c *Command) decodeGetCommand(msg protocol.Command) error {
	getCmd, err := msg.SubCommand().Get()
	if err != nil {
		return err
	}

	key, err := getCmd.Key()
	if err != nil {
		return err
	}

	c.SubCommand = NewGetCommand(key)
	c.Action = c.SubCommand.getAction()
	return nil
}

func (c *Command) decodeHandshakeCommand(msg protocol.Command) error {
	subCmd, err := msg.SubCommand().Handshake()
	if err != nil {
		return err
	}

	userAgent, err := subCmd.UserAgent()
	if err != nil {
		return err
	}

	c.SubCommand = NewHandshakeCommand(userAgent)
	c.Action = c.SubCommand.getAction()
	return nil
}
