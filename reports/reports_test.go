package reports

import (
	"bookexchange/model"
	"testing"
)

func TestSummary(t *testing.T) {
	s := BuildSummary([]model.Record{{ID: "1", OwnerID: "u", Status: "available"}, {ID: "2", OwnerID: "u", Status: "archived"}})
	if s.Total != 2 || s.Available != 1 {
		t.Fatal(s)
	}
}
