package store

import (
	"bookexchange/model"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := model.Record{ID: "r1", Title: "Go", OwnerID: "u1"}
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.LoadRecord("r1")
	if e != nil || got.Title != "Go" {
		t.Fatalf("%v %#v", e, got)
	}
}
