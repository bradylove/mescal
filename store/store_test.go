package store_test

import (
	"bytes"
	"github.com/bradylove/mescal/msg"
	"github.com/bradylove/mescal/store"
	"testing"
)

const (
	succeed = "\033[0;32m\u2713\033[0m"
	failed  = "\033[0;31m\u2717\033[0m"
)

func TestHandleCommandWithGetCommand(t *testing.T) {
	t.Log("Given I have a get command")
	{
		var b []byte
		buf := bytes.NewBuffer(b)
		sr := store.NewStore()
		cmd := msg.NewCommand("12345", msg.NewGetCommand("foo"))
		sr.HandleCommand(cmd, buf)
		sr.Close()
		sr.Wait()

		if len(buf.Bytes()) > 40 {
			t.Logf("It should have bytes %s", succeed)
		} else {
			t.Fatalf("It should have bytes %s", failed)
		}
	}
}

//func TestHandleCommandWithSetCommand(t *testing.T) {
//	t.Log("Given I have a set command")
//	{
//	}
//}
