package bookexchange

import (
	"bookexchange/model"
	"bookexchange/store"
	"bookexchange/workflow"
	"path/filepath"
	"testing"
)

func TestWorkflowThree(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	w := workflow.New(s)
	if e := w.Receive(model.Record{ID: "r", Title: "History", OwnerID: "u"}); e != nil {
		t.Fatal(e)
	}
	if e := w.Process("r", "u"); e != nil {
		t.Fatal(e)
	}
	if e := w.Exchange("r", "u"); e != nil {
		t.Fatal(e)
	}
	if got := w.QueryRecords("history").Total; got != 1 {
		t.Fatal(got)
	}
}
