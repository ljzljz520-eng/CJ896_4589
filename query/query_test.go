package query

import (
	"bookexchange/catalog"
	"bookexchange/model"
	"bookexchange/store"
	"path/filepath"
	"testing"
)

func TestQueries(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := catalog.New(s)
	_ = c.Register(model.Record{ID: "r", Title: "Book", Author: "A", OwnerID: "u", Status: "available"})
	q := New(c)
	if len(q.Available()) != 1 {
		t.Fatal("not available")
	}
}
