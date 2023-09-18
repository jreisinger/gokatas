// Package kv implements key-value store that persists data to disk. The process
// of sending data out of the program is called marshalling. We could marshal Go
// data into many formats: plain text, base64-encoded text, SQL results or
// queries, TCP/IP packets, ... The simplest way to format data for transmission
// is a stream of bytes. This kind of marshalling is called serialization. The
// encoding/gob package can serialise most kinds of Go values to bytes. Adapted
// from: https://github.com/bitfield/tpg-tools2/tree/main/kv
//
// Level: Intermediate
// Topics: encoding/gob, marshaling, tpg-tools
package kv

import (
	"encoding/gob"
	"errors"
	"io/fs"
	"os"
)

type Store struct {
	path string
	data map[string]string
}

func OpenStore(path string) (*Store, error) {
	s := &Store{
		path: path,
		data: make(map[string]string),
	}
	f, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
	if err = gob.NewDecoder(f).Decode(&s.data); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Store) Save() error {
	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer f.Close()
	return gob.NewEncoder(f).Encode(s.data)
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}

func (s *Store) Get(key string) (string, bool) {
	v, ok := s.data[key]
	return v, ok
}
