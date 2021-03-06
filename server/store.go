package main

import (
	"fmt"
	"github.com/bradylove/mescal/msg"
	"io"
	"sync"
)

type object struct {
	key    string
	value  string
	expiry int64
}

type objects map[string]object

type message struct {
	*msg.Command

	writer io.Writer
}

type Store struct {
	objects objects
	wg      sync.WaitGroup
	msgs    chan message
}

const (
	workerCount = 20
)

func NewStore() *Store {
	s := Store{
		objects: objects{},
		msgs:    make(chan message),
	}

	s.setupPool()

	return &s
}

func (s *Store) setupPool() {
	s.wg.Add(workerCount)

	// TODO: Make this smarter, each connection should get its own go routine for
	// talking to the store
	for i := 0; i < workerCount; i++ {
		go s.manageStore()
	}
}

func (s *Store) manageStore() {
StoreHandler:
	for {
		m, ok := <-s.msgs
		if !ok {
			break StoreHandler
		}

		switch sb := m.SubCommand.(type) {
		case msg.GetCommand:
			s.handleGetCommand(m.Command, sb, m.writer)
		case msg.SetCommand:
			s.handleSetCommand(m.Command, sb, m.writer)
		default:
			fmt.Println("Unknown command")
		}
	}

	s.wg.Done()
}

func (s *Store) handleGetCommand(cmd *msg.Command, sb msg.GetCommand, w io.Writer) {
	obj := s.objects[sb.Key]

	res := msg.NewResult(
		cmd.Id,
		msg.StatusSuccess,
		msg.NewGetResult(obj.key, obj.value, obj.expiry),
	)

	res.Encode(w)
}

func (s *Store) handleSetCommand(cmd *msg.Command, sb msg.SetCommand, w io.Writer) {
	s.objects[sb.Key] = object{sb.Key, sb.Value, sb.Expiry}

	res := msg.NewResult(
		cmd.Id,
		msg.StatusSuccess,
		msg.NewSetResult(),
	)

	res.Encode(w)
}

func (s *Store) HandleCommand(cmd *msg.Command, w io.Writer) {
	s.msgs <- message{cmd, w}
}

func (s *Store) wait() {
	s.wg.Wait()
}

func (s *Store) close() {
	close(s.msgs)
}
