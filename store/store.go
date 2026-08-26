package store

import (
	"bookexchange/model"
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"sync"
)

var buckets = [][]byte{[]byte("records"), []byte("profiles"), []byte("events"), []byte("audits"), []byte("listings"), []byte("notifications")}

type Store struct {
	db   *bbolt.DB
	mu   sync.RWMutex
	path string
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db, path: path}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	e := s.db.Close()
	s.db = nil
	return e
}
func (s *Store) Put(bucket, key string, v any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), data) })
}
func (s *Store) Get(bucket, key string, out any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if v == nil {
			return bbolt.ErrBucketNotFound
		}
		return json.Unmarshal(v, out)
	})
}
func (s *Store) Delete(bucket, key string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Delete([]byte(key)) })
}
func (s *Store) List(bucket string) ([][]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var out [][]byte
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).ForEach(func(k, v []byte) error { out = append(out, append([]byte(nil), v...)); return nil })
	})
	return out, e
}
func (s *Store) SaveRecord(r model.Record) error { return s.Put("records", r.ID, r) }
func (s *Store) LoadRecord(id string) (model.Record, error) {
	var r model.Record
	e := s.Get("records", id, &r)
	return r, e
}
