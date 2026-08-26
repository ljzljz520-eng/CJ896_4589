package catalog

import (
	"bookexchange/model"
	"testing"
)

func TestBusinessChain05(t *testing.T) {
	entries := []model.Record{{ID: "a", Title: "A", OwnerID: "u"}, {ID: "b", Title: "B", OwnerID: "u"}}
	old := CloneEntries(entries)
	entries[0].Title = "changed"
	if old[0].Title != "A" {
		t.Fatal("previous listing changed")
	}
}
