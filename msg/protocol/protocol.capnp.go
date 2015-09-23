package protocol

// AUTO GENERATED - DO NOT EDIT

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto"
)

type Action uint16

// Values of Action.
const (
	Action_handshake Action = 0
	Action_get       Action = 1
	Action_set       Action = 2
)

// String returns the enum's constant name.
func (c Action) String() string {
	switch c {
	case Action_handshake:
		return "handshake"
	case Action_get:
		return "get"
	case Action_set:
		return "set"

	default:
		return ""
	}
}

// ActionFromString returns the enum value with a name,
// or the zero value if there's no such value.
func ActionFromString(c string) Action {
	switch c {
	case "handshake":
		return Action_handshake
	case "get":
		return Action_get
	case "set":
		return Action_set

	default:
		return 0
	}
}

type Action_List struct{ capnp.List }

func NewAction_List(s *capnp.Segment, sz int32) (Action_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	if err != nil {
		return Action_List{}, err
	}
	return Action_List{l.List}, nil
}

func (l Action_List) At(i int) Action {
	ul := capnp.UInt16List{List: l.List}
	return Action(ul.At(i))
}

func (l Action_List) Set(i int, v Action) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type Status uint16

// Values of Status.
const (
	Status_success Status = 0
)

// String returns the enum's constant name.
func (c Status) String() string {
	switch c {
	case Status_success:
		return "success"

	default:
		return ""
	}
}

// StatusFromString returns the enum value with a name,
// or the zero value if there's no such value.
func StatusFromString(c string) Status {
	switch c {
	case "success":
		return Status_success

	default:
		return 0
	}
}

type Status_List struct{ capnp.List }

func NewStatus_List(s *capnp.Segment, sz int32) (Status_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	if err != nil {
		return Status_List{}, err
	}
	return Status_List{l.List}, nil
}

func (l Status_List) At(i int) Status {
	ul := capnp.UInt16List{List: l.List}
	return Status(ul.At(i))
}

func (l Status_List) Set(i int, v Status) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type Command struct{ capnp.Struct }
type Command_subCommand Command
type Command_subCommand_Which uint16

const (
	Command_subCommand_Which_handshake Command_subCommand_Which = 0
	Command_subCommand_Which_get       Command_subCommand_Which = 1
	Command_subCommand_Which_set       Command_subCommand_Which = 2
)

func (w Command_subCommand_Which) String() string {
	const s = "handshakegetset"
	switch w {
	case Command_subCommand_Which_handshake:
		return s[0:9]
	case Command_subCommand_Which_get:
		return s[9:12]
	case Command_subCommand_Which_set:
		return s[12:15]

	}
	return "Command_subCommand_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewCommand(s *capnp.Segment) (Command, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Command{}, err
	}
	return Command{st}, nil
}

func NewRootCommand(s *capnp.Segment) (Command, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Command{}, err
	}
	return Command{st}, nil
}

func ReadRootCommand(msg *capnp.Message) (Command, error) {
	root, err := msg.Root()
	if err != nil {
		return Command{}, err
	}
	st := capnp.ToStruct(root)
	return Command{st}, nil
}

func (s Command) Id() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Command) SetId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s Command) Action() Action {
	return Action(s.Struct.Uint16(0))
}

func (s Command) SetAction(v Action) {

	s.Struct.SetUint16(0, uint16(v))
}
func (s Command) SubCommand() Command_subCommand { return Command_subCommand(s) }

func (s Command_subCommand) Which() Command_subCommand_Which {
	return Command_subCommand_Which(s.Struct.Uint16(2))
}

func (s Command_subCommand) Handshake() (Command_HandshakeCommand, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return Command_HandshakeCommand{}, err
	}

	ss := capnp.ToStruct(p)

	return Command_HandshakeCommand{Struct: ss}, nil
}

func (s Command_subCommand) SetHandshake(v Command_HandshakeCommand) error {
	s.Struct.SetUint16(2, 0)
	return s.Struct.SetPointer(1, v.Struct)
}

// NewHandshake sets the handshake field to a newly
// allocated Command_HandshakeCommand struct, preferring placement in s's segment.
func (s Command_subCommand) NewHandshake() (Command_HandshakeCommand, error) {
	s.Struct.SetUint16(2, 0)
	ss, err := NewCommand_HandshakeCommand(s.Struct.Segment())
	if err != nil {
		return Command_HandshakeCommand{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

func (s Command_subCommand) Get() (Command_GetCommand, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return Command_GetCommand{}, err
	}

	ss := capnp.ToStruct(p)

	return Command_GetCommand{Struct: ss}, nil
}

func (s Command_subCommand) SetGet(v Command_GetCommand) error {
	s.Struct.SetUint16(2, 1)
	return s.Struct.SetPointer(1, v.Struct)
}

// NewGet sets the get field to a newly
// allocated Command_GetCommand struct, preferring placement in s's segment.
func (s Command_subCommand) NewGet() (Command_GetCommand, error) {
	s.Struct.SetUint16(2, 1)
	ss, err := NewCommand_GetCommand(s.Struct.Segment())
	if err != nil {
		return Command_GetCommand{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

func (s Command_subCommand) Set() (Command_SetCommand, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return Command_SetCommand{}, err
	}

	ss := capnp.ToStruct(p)

	return Command_SetCommand{Struct: ss}, nil
}

func (s Command_subCommand) SetSet(v Command_SetCommand) error {
	s.Struct.SetUint16(2, 2)
	return s.Struct.SetPointer(1, v.Struct)
}

// NewSet sets the set field to a newly
// allocated Command_SetCommand struct, preferring placement in s's segment.
func (s Command_subCommand) NewSet() (Command_SetCommand, error) {
	s.Struct.SetUint16(2, 2)
	ss, err := NewCommand_SetCommand(s.Struct.Segment())
	if err != nil {
		return Command_SetCommand{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

// Command_List is a list of Command.
type Command_List struct{ capnp.List }

// NewCommand creates a new list of Command.
func NewCommand_List(s *capnp.Segment, sz int32) (Command_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	if err != nil {
		return Command_List{}, err
	}
	return Command_List{l}, nil
}

func (s Command_List) At(i int) Command           { return Command{s.List.Struct(i)} }
func (s Command_List) Set(i int, v Command) error { return s.List.SetStruct(i, v.Struct) }

// Command_Promise is a wrapper for a Command promised by a client call.
type Command_Promise struct{ *capnp.Pipeline }

func (p Command_Promise) Struct() (Command, error) {
	s, err := p.Pipeline.Struct()
	return Command{s}, err
}
func (p Command_Promise) SubCommand() Command_subCommand_Promise {
	return Command_subCommand_Promise{p.Pipeline}
}

// Command_subCommand_Promise is a wrapper for a Command_subCommand promised by a client call.
type Command_subCommand_Promise struct{ *capnp.Pipeline }

func (p Command_subCommand_Promise) Struct() (Command_subCommand, error) {
	s, err := p.Pipeline.Struct()
	return Command_subCommand{s}, err
}

func (p Command_subCommand_Promise) Handshake() Command_HandshakeCommand_Promise {
	return Command_HandshakeCommand_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Command_subCommand_Promise) Get() Command_GetCommand_Promise {
	return Command_GetCommand_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Command_subCommand_Promise) Set() Command_SetCommand_Promise {
	return Command_SetCommand_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type Command_HandshakeCommand struct{ capnp.Struct }

func NewCommand_HandshakeCommand(s *capnp.Segment) (Command_HandshakeCommand, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Command_HandshakeCommand{}, err
	}
	return Command_HandshakeCommand{st}, nil
}

func NewRootCommand_HandshakeCommand(s *capnp.Segment) (Command_HandshakeCommand, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Command_HandshakeCommand{}, err
	}
	return Command_HandshakeCommand{st}, nil
}

func ReadRootCommand_HandshakeCommand(msg *capnp.Message) (Command_HandshakeCommand, error) {
	root, err := msg.Root()
	if err != nil {
		return Command_HandshakeCommand{}, err
	}
	st := capnp.ToStruct(root)
	return Command_HandshakeCommand{st}, nil
}

func (s Command_HandshakeCommand) UserAgent() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Command_HandshakeCommand) SetUserAgent(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// Command_HandshakeCommand_List is a list of Command_HandshakeCommand.
type Command_HandshakeCommand_List struct{ capnp.List }

// NewCommand_HandshakeCommand creates a new list of Command_HandshakeCommand.
func NewCommand_HandshakeCommand_List(s *capnp.Segment, sz int32) (Command_HandshakeCommand_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Command_HandshakeCommand_List{}, err
	}
	return Command_HandshakeCommand_List{l}, nil
}

func (s Command_HandshakeCommand_List) At(i int) Command_HandshakeCommand {
	return Command_HandshakeCommand{s.List.Struct(i)}
}
func (s Command_HandshakeCommand_List) Set(i int, v Command_HandshakeCommand) error {
	return s.List.SetStruct(i, v.Struct)
}

// Command_HandshakeCommand_Promise is a wrapper for a Command_HandshakeCommand promised by a client call.
type Command_HandshakeCommand_Promise struct{ *capnp.Pipeline }

func (p Command_HandshakeCommand_Promise) Struct() (Command_HandshakeCommand, error) {
	s, err := p.Pipeline.Struct()
	return Command_HandshakeCommand{s}, err
}

type Command_GetCommand struct{ capnp.Struct }

func NewCommand_GetCommand(s *capnp.Segment) (Command_GetCommand, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Command_GetCommand{}, err
	}
	return Command_GetCommand{st}, nil
}

func NewRootCommand_GetCommand(s *capnp.Segment) (Command_GetCommand, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Command_GetCommand{}, err
	}
	return Command_GetCommand{st}, nil
}

func ReadRootCommand_GetCommand(msg *capnp.Message) (Command_GetCommand, error) {
	root, err := msg.Root()
	if err != nil {
		return Command_GetCommand{}, err
	}
	st := capnp.ToStruct(root)
	return Command_GetCommand{st}, nil
}

func (s Command_GetCommand) Key() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Command_GetCommand) SetKey(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// Command_GetCommand_List is a list of Command_GetCommand.
type Command_GetCommand_List struct{ capnp.List }

// NewCommand_GetCommand creates a new list of Command_GetCommand.
func NewCommand_GetCommand_List(s *capnp.Segment, sz int32) (Command_GetCommand_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Command_GetCommand_List{}, err
	}
	return Command_GetCommand_List{l}, nil
}

func (s Command_GetCommand_List) At(i int) Command_GetCommand {
	return Command_GetCommand{s.List.Struct(i)}
}
func (s Command_GetCommand_List) Set(i int, v Command_GetCommand) error {
	return s.List.SetStruct(i, v.Struct)
}

// Command_GetCommand_Promise is a wrapper for a Command_GetCommand promised by a client call.
type Command_GetCommand_Promise struct{ *capnp.Pipeline }

func (p Command_GetCommand_Promise) Struct() (Command_GetCommand, error) {
	s, err := p.Pipeline.Struct()
	return Command_GetCommand{s}, err
}

type Command_SetCommand struct{ capnp.Struct }

func NewCommand_SetCommand(s *capnp.Segment) (Command_SetCommand, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Command_SetCommand{}, err
	}
	return Command_SetCommand{st}, nil
}

func NewRootCommand_SetCommand(s *capnp.Segment) (Command_SetCommand, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Command_SetCommand{}, err
	}
	return Command_SetCommand{st}, nil
}

func ReadRootCommand_SetCommand(msg *capnp.Message) (Command_SetCommand, error) {
	root, err := msg.Root()
	if err != nil {
		return Command_SetCommand{}, err
	}
	st := capnp.ToStruct(root)
	return Command_SetCommand{st}, nil
}

func (s Command_SetCommand) Key() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Command_SetCommand) SetKey(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s Command_SetCommand) Value() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Command_SetCommand) SetValue(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s Command_SetCommand) Expiry() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Command_SetCommand) SetExpiry(v int64) {

	s.Struct.SetUint64(0, uint64(v))
}

// Command_SetCommand_List is a list of Command_SetCommand.
type Command_SetCommand_List struct{ capnp.List }

// NewCommand_SetCommand creates a new list of Command_SetCommand.
func NewCommand_SetCommand_List(s *capnp.Segment, sz int32) (Command_SetCommand_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	if err != nil {
		return Command_SetCommand_List{}, err
	}
	return Command_SetCommand_List{l}, nil
}

func (s Command_SetCommand_List) At(i int) Command_SetCommand {
	return Command_SetCommand{s.List.Struct(i)}
}
func (s Command_SetCommand_List) Set(i int, v Command_SetCommand) error {
	return s.List.SetStruct(i, v.Struct)
}

// Command_SetCommand_Promise is a wrapper for a Command_SetCommand promised by a client call.
type Command_SetCommand_Promise struct{ *capnp.Pipeline }

func (p Command_SetCommand_Promise) Struct() (Command_SetCommand, error) {
	s, err := p.Pipeline.Struct()
	return Command_SetCommand{s}, err
}

type Result struct{ capnp.Struct }
type Result_subResult Result
type Result_subResult_Which uint16

const (
	Result_subResult_Which_handshake Result_subResult_Which = 0
	Result_subResult_Which_get       Result_subResult_Which = 1
	Result_subResult_Which_set       Result_subResult_Which = 2
)

func (w Result_subResult_Which) String() string {
	const s = "handshakegetset"
	switch w {
	case Result_subResult_Which_handshake:
		return s[0:9]
	case Result_subResult_Which_get:
		return s[9:12]
	case Result_subResult_Which_set:
		return s[12:15]

	}
	return "Result_subResult_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewResult(s *capnp.Segment) (Result, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Result{}, err
	}
	return Result{st}, nil
}

func NewRootResult(s *capnp.Segment) (Result, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Result{}, err
	}
	return Result{st}, nil
}

func ReadRootResult(msg *capnp.Message) (Result, error) {
	root, err := msg.Root()
	if err != nil {
		return Result{}, err
	}
	st := capnp.ToStruct(root)
	return Result{st}, nil
}

func (s Result) Id() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Result) SetId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s Result) Action() Action {
	return Action(s.Struct.Uint16(0))
}

func (s Result) SetAction(v Action) {

	s.Struct.SetUint16(0, uint16(v))
}

func (s Result) Status() Status {
	return Status(s.Struct.Uint16(2))
}

func (s Result) SetStatus(v Status) {

	s.Struct.SetUint16(2, uint16(v))
}
func (s Result) SubResult() Result_subResult { return Result_subResult(s) }

func (s Result_subResult) Which() Result_subResult_Which {
	return Result_subResult_Which(s.Struct.Uint16(4))
}

func (s Result_subResult) Handshake() (Result_HandshakeResult, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return Result_HandshakeResult{}, err
	}

	ss := capnp.ToStruct(p)

	return Result_HandshakeResult{Struct: ss}, nil
}

func (s Result_subResult) SetHandshake(v Result_HandshakeResult) error {
	s.Struct.SetUint16(4, 0)
	return s.Struct.SetPointer(1, v.Struct)
}

// NewHandshake sets the handshake field to a newly
// allocated Result_HandshakeResult struct, preferring placement in s's segment.
func (s Result_subResult) NewHandshake() (Result_HandshakeResult, error) {
	s.Struct.SetUint16(4, 0)
	ss, err := NewResult_HandshakeResult(s.Struct.Segment())
	if err != nil {
		return Result_HandshakeResult{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

func (s Result_subResult) Get() (Result_GetResult, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return Result_GetResult{}, err
	}

	ss := capnp.ToStruct(p)

	return Result_GetResult{Struct: ss}, nil
}

func (s Result_subResult) SetGet(v Result_GetResult) error {
	s.Struct.SetUint16(4, 1)
	return s.Struct.SetPointer(1, v.Struct)
}

// NewGet sets the get field to a newly
// allocated Result_GetResult struct, preferring placement in s's segment.
func (s Result_subResult) NewGet() (Result_GetResult, error) {
	s.Struct.SetUint16(4, 1)
	ss, err := NewResult_GetResult(s.Struct.Segment())
	if err != nil {
		return Result_GetResult{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

func (s Result_subResult) Set() (Result_SetResult, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return Result_SetResult{}, err
	}

	ss := capnp.ToStruct(p)

	return Result_SetResult{Struct: ss}, nil
}

func (s Result_subResult) SetSet(v Result_SetResult) error {
	s.Struct.SetUint16(4, 2)
	return s.Struct.SetPointer(1, v.Struct)
}

// NewSet sets the set field to a newly
// allocated Result_SetResult struct, preferring placement in s's segment.
func (s Result_subResult) NewSet() (Result_SetResult, error) {
	s.Struct.SetUint16(4, 2)
	ss, err := NewResult_SetResult(s.Struct.Segment())
	if err != nil {
		return Result_SetResult{}, err
	}
	err = s.Struct.SetPointer(1, ss)
	return ss, err
}

// Result_List is a list of Result.
type Result_List struct{ capnp.List }

// NewResult creates a new list of Result.
func NewResult_List(s *capnp.Segment, sz int32) (Result_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	if err != nil {
		return Result_List{}, err
	}
	return Result_List{l}, nil
}

func (s Result_List) At(i int) Result           { return Result{s.List.Struct(i)} }
func (s Result_List) Set(i int, v Result) error { return s.List.SetStruct(i, v.Struct) }

// Result_Promise is a wrapper for a Result promised by a client call.
type Result_Promise struct{ *capnp.Pipeline }

func (p Result_Promise) Struct() (Result, error) {
	s, err := p.Pipeline.Struct()
	return Result{s}, err
}
func (p Result_Promise) SubResult() Result_subResult_Promise {
	return Result_subResult_Promise{p.Pipeline}
}

// Result_subResult_Promise is a wrapper for a Result_subResult promised by a client call.
type Result_subResult_Promise struct{ *capnp.Pipeline }

func (p Result_subResult_Promise) Struct() (Result_subResult, error) {
	s, err := p.Pipeline.Struct()
	return Result_subResult{s}, err
}

func (p Result_subResult_Promise) Handshake() Result_HandshakeResult_Promise {
	return Result_HandshakeResult_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Result_subResult_Promise) Get() Result_GetResult_Promise {
	return Result_GetResult_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Result_subResult_Promise) Set() Result_SetResult_Promise {
	return Result_SetResult_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type Result_HandshakeResult struct{ capnp.Struct }

func NewResult_HandshakeResult(s *capnp.Segment) (Result_HandshakeResult, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Result_HandshakeResult{}, err
	}
	return Result_HandshakeResult{st}, nil
}

func NewRootResult_HandshakeResult(s *capnp.Segment) (Result_HandshakeResult, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Result_HandshakeResult{}, err
	}
	return Result_HandshakeResult{st}, nil
}

func ReadRootResult_HandshakeResult(msg *capnp.Message) (Result_HandshakeResult, error) {
	root, err := msg.Root()
	if err != nil {
		return Result_HandshakeResult{}, err
	}
	st := capnp.ToStruct(root)
	return Result_HandshakeResult{st}, nil
}

func (s Result_HandshakeResult) ClientId() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Result_HandshakeResult) SetClientId(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

// Result_HandshakeResult_List is a list of Result_HandshakeResult.
type Result_HandshakeResult_List struct{ capnp.List }

// NewResult_HandshakeResult creates a new list of Result_HandshakeResult.
func NewResult_HandshakeResult_List(s *capnp.Segment, sz int32) (Result_HandshakeResult_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Result_HandshakeResult_List{}, err
	}
	return Result_HandshakeResult_List{l}, nil
}

func (s Result_HandshakeResult_List) At(i int) Result_HandshakeResult {
	return Result_HandshakeResult{s.List.Struct(i)}
}
func (s Result_HandshakeResult_List) Set(i int, v Result_HandshakeResult) error {
	return s.List.SetStruct(i, v.Struct)
}

// Result_HandshakeResult_Promise is a wrapper for a Result_HandshakeResult promised by a client call.
type Result_HandshakeResult_Promise struct{ *capnp.Pipeline }

func (p Result_HandshakeResult_Promise) Struct() (Result_HandshakeResult, error) {
	s, err := p.Pipeline.Struct()
	return Result_HandshakeResult{s}, err
}

type Result_GetResult struct{ capnp.Struct }

func NewResult_GetResult(s *capnp.Segment) (Result_GetResult, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Result_GetResult{}, err
	}
	return Result_GetResult{st}, nil
}

func NewRootResult_GetResult(s *capnp.Segment) (Result_GetResult, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	if err != nil {
		return Result_GetResult{}, err
	}
	return Result_GetResult{st}, nil
}

func ReadRootResult_GetResult(msg *capnp.Message) (Result_GetResult, error) {
	root, err := msg.Root()
	if err != nil {
		return Result_GetResult{}, err
	}
	st := capnp.ToStruct(root)
	return Result_GetResult{st}, nil
}

func (s Result_GetResult) Key() (string, error) {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Result_GetResult) SetKey(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(0, t)
}

func (s Result_GetResult) Value() (string, error) {
	p, err := s.Struct.Pointer(1)
	if err != nil {
		return "", err
	}

	return capnp.ToText(p), nil

}

func (s Result_GetResult) SetValue(v string) error {

	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPointer(1, t)
}

func (s Result_GetResult) Expiry() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Result_GetResult) SetExpiry(v int64) {

	s.Struct.SetUint64(0, uint64(v))
}

// Result_GetResult_List is a list of Result_GetResult.
type Result_GetResult_List struct{ capnp.List }

// NewResult_GetResult creates a new list of Result_GetResult.
func NewResult_GetResult_List(s *capnp.Segment, sz int32) (Result_GetResult_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	if err != nil {
		return Result_GetResult_List{}, err
	}
	return Result_GetResult_List{l}, nil
}

func (s Result_GetResult_List) At(i int) Result_GetResult { return Result_GetResult{s.List.Struct(i)} }
func (s Result_GetResult_List) Set(i int, v Result_GetResult) error {
	return s.List.SetStruct(i, v.Struct)
}

// Result_GetResult_Promise is a wrapper for a Result_GetResult promised by a client call.
type Result_GetResult_Promise struct{ *capnp.Pipeline }

func (p Result_GetResult_Promise) Struct() (Result_GetResult, error) {
	s, err := p.Pipeline.Struct()
	return Result_GetResult{s}, err
}

type Result_SetResult struct{ capnp.Struct }

func NewResult_SetResult(s *capnp.Segment) (Result_SetResult, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Result_SetResult{}, err
	}
	return Result_SetResult{st}, nil
}

func NewRootResult_SetResult(s *capnp.Segment) (Result_SetResult, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Result_SetResult{}, err
	}
	return Result_SetResult{st}, nil
}

func ReadRootResult_SetResult(msg *capnp.Message) (Result_SetResult, error) {
	root, err := msg.Root()
	if err != nil {
		return Result_SetResult{}, err
	}
	st := capnp.ToStruct(root)
	return Result_SetResult{st}, nil
}

// Result_SetResult_List is a list of Result_SetResult.
type Result_SetResult_List struct{ capnp.List }

// NewResult_SetResult creates a new list of Result_SetResult.
func NewResult_SetResult_List(s *capnp.Segment, sz int32) (Result_SetResult_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Result_SetResult_List{}, err
	}
	return Result_SetResult_List{l}, nil
}

func (s Result_SetResult_List) At(i int) Result_SetResult { return Result_SetResult{s.List.Struct(i)} }
func (s Result_SetResult_List) Set(i int, v Result_SetResult) error {
	return s.List.SetStruct(i, v.Struct)
}

// Result_SetResult_Promise is a wrapper for a Result_SetResult promised by a client call.
type Result_SetResult_Promise struct{ *capnp.Pipeline }

func (p Result_SetResult_Promise) Struct() (Result_SetResult, error) {
	s, err := p.Pipeline.Struct()
	return Result_SetResult{s}, err
}
