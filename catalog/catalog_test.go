package catalog

import (
	"bookexchange/model"
	"bookexchange/store"
	"path/filepath"
	"testing"
)

func TestRegisterAndSearch(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := New(s)
	_ = c.Register(model.Record{ID: "1", Title: "Clean Code", Author: "Martin", OwnerID: "u", Tags: []string{"Go"}})
	if c.Search("clean").Total != 1 {
		t.Fatal("missing")
	}
}
