package store

import "github.com/bradylove/mescal/pkg/mescal"

type Store struct {
	cache map[string]*mescal.KeyValue
}

func New() *Store {
	return &Store{cache: make(map[string]*mescal.KeyValue)}
}

func (s Store) Set(kv *mescal.KeyValue) error {
	s.cache[kv.Key] = kv

	return nil
}

func (s Store) Get(key string) (*mescal.KeyValue, error) {
	return s.cache[key], nil
}
