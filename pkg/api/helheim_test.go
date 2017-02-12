// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package api_test

import "github.com/bradylove/mescal/pkg/mescal"

type mockStore struct {
	SetCalled chan bool
	SetInput  struct {
		KeyValue chan *mescal.KeyValue
	}
	SetOutput struct {
		Err chan error
	}
	GetCalled chan bool
	GetInput  struct {
		Key chan string
	}
	GetOutput struct {
		KeyValue chan *mescal.KeyValue
		Err      chan error
	}
}

func newMockStore() *mockStore {
	m := &mockStore{}
	m.SetCalled = make(chan bool, 100)
	m.SetInput.KeyValue = make(chan *mescal.KeyValue, 100)
	m.SetOutput.Err = make(chan error, 100)
	m.GetCalled = make(chan bool, 100)
	m.GetInput.Key = make(chan string, 100)
	m.GetOutput.KeyValue = make(chan *mescal.KeyValue, 100)
	m.GetOutput.Err = make(chan error, 100)
	return m
}
func (m *mockStore) Set(keyValue *mescal.KeyValue) (err error) {
	m.SetCalled <- true
	m.SetInput.KeyValue <- keyValue
	return <-m.SetOutput.Err
}
func (m *mockStore) Get(key string) (keyValue *mescal.KeyValue, err error) {
	m.GetCalled <- true
	m.GetInput.Key <- key
	return <-m.GetOutput.KeyValue, <-m.GetOutput.Err
}