using Go = import "../../../../../zombiezen.com/go/capnproto/go.capnp";
@0xd678d303d1e1ea2c;
$Go.package("protocol");
$Go.import("protocol");

enum Action {
  handshake @0;
  get @1;
  set @2;
}

enum Status {
  success @0;
  clientError @1
  serverError @2
  keyNotFound @3
}

struct Command {
  id @0 :Text;
  action @1 :Action;

  subCommand :union {
    handshake @2 :HandshakeCommand;
    get @3 :GetCommand;
    set @4 :SetCommand;
  }

  struct HandshakeCommand {
    userAgent @0 :Text;
  }

  struct GetCommand {
    key @0 :Text;
  }

  struct SetCommand {
    key @0 :Text;
    value @1 :Text;
    expiry @2 :Int64;
  }
}

struct Result {
  id @0 :Text;
  action @1 :Action;
  status @2 :Status;

  subResult :union {
    handshake @3 :HandshakeResult;
    get @4 :GetResult;
    set @5 :SetResult;
  }

  struct HandshakeResult {
    clientId @0 :Text;
  }

  struct GetResult {
    key @0 :Text;
    value @1 :Text;
    expiry @2 :Int64;
  }

  struct SetResult {
  }
}
