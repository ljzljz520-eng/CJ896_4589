package archive

import (
	"bookexchange/catalog"
	"bookexchange/model"
	"bookexchange/store"
	"path/filepath"
	"testing"
)

func TestArchiveLifecycle(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := catalog.New(s)
	a := New(s)
	_ = c.Register(model.Record{ID: "r", Title: "B", OwnerID: "u"})
	if e := a.StoreRecord("r", c); e != nil {
		t.Fatal(e)
	}
	if e := a.Restore("r", c); e != nil {
		t.Fatal(e)
	}
}
