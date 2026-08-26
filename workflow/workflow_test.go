package workflow

import (
	"bookexchange/model"
	"bookexchange/store"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	w := New(s)
	if e := w.Receive(model.Record{ID: "r", Title: "Book", OwnerID: "u"}); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	w := New(s)
	_ = w.Receive(model.Record{ID: "r", Title: "Book", OwnerID: "u"})
	if e := w.Process("r", "u"); e != nil {
		t.Fatal(e)
	}
}
