package store

import (
	"errors"
	"hash/crc64"
	"log"
	"sync/atomic"
	"unsafe"

	"github.com/bradylove/mescal/pkg/mescal"
)

var (
	bucketCount     = uint64(1000003)
	sliceCount      = uint64(53)
	allowableOffset = uint64(5)
	tablePoly       = uint64(23814672813913)

	ErrKeyNotFound = errors.New("Key not found")
)

type Store struct {
	cache     [][]unsafe.Pointer
	polyTable *crc64.Table
}

func New() *Store {
	s := Store{
		cache:     make([][]unsafe.Pointer, sliceCount),
		polyTable: crc64.MakeTable(tablePoly),
	}

	for idx := range s.cache {
		s.cache[uint64(idx)] = make([]unsafe.Pointer, bucketCount)
	}

	return &s
}

func (s *Store) Set(kv *mescal.KeyValue) error {
	slIdx, bIdx := s.buildIndex(kv.Key)

	// Allow the value to be placed in correct bucket or 4 buckets to the right.
	for i := uint64(0); i < allowableOffset; i++ {
		if s.compareAndSwap(slIdx, (bIdx+i)%bucketCount, kv) {
			return nil
		}
	}

	// Allow the value to be placed in correct slice or 4 slices down.
	for i := uint64(0); i < allowableOffset; i++ {
		if s.compareAndSwap((slIdx+i)%sliceCount, bIdx, kv) {
			return nil
		}
	}

	// All calculations ended up in a collision, place the key/value in the
	// originally calculated location
	log.Println("Failed to find open bucket for key/value, placing in original bucket")
	atomic.SwapPointer(
		&s.cache[slIdx][bIdx],
		unsafe.Pointer(kv),
	)

	return nil
}

func (s *Store) Get(key string) (*mescal.KeyValue, error) {
	slIdx, bIdx := s.buildIndex(key)

	// Look for the key in the correct bucket, if the key does not match try
	// the next four buckets to the right
	for i := uint64(0); i < allowableOffset; i++ {
		if kv := s.findKeyValueAtIndex(slIdx, (bIdx+i)%bucketCount, key); kv != nil {
			return kv, nil
		}
	}

	// Look for the key in the correct bucket, if the key does not match try
	// the next four slices down
	for i := uint64(0); i < allowableOffset; i++ {
		if kv := s.findKeyValueAtIndex((slIdx+i)%bucketCount, bIdx, key); kv != nil {
			return kv, nil
		}
	}

	return nil, ErrKeyNotFound
}

func (s *Store) buildIndex(key string) (sliceIdx uint64, bucketIdx uint64) {
	f := crc64.New(s.polyTable)
	f.Write([]byte(key))

	sum := f.Sum64()

	return sum % sliceCount, sum % bucketCount
}

func (s *Store) compareAndSwap(sliceIdx, bucketIdx uint64, kv *mescal.KeyValue) bool {
	old := (*mescal.KeyValue)(atomic.LoadPointer(&s.cache[sliceIdx][bucketIdx]))

	if old == nil || old.Key == kv.Key {
		atomic.SwapPointer(
			&s.cache[sliceIdx][bucketIdx],
			unsafe.Pointer(kv),
		)

		return true
	}

	return false
}

func (s *Store) findKeyValueAtIndex(sliceIdx, bucketIdx uint64, key string) *mescal.KeyValue {
	kv := (*mescal.KeyValue)(atomic.LoadPointer(&s.cache[sliceIdx][bucketIdx]))

	if kv != nil && kv.Key == key {
		return kv
	}

	return nil
}
