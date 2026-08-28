package catalog

import (
	"bookexchange/model"
	"strings"
	"time"
)

type Metrics struct {
	Seen, Updated, Rejected int
	LastUpdate              time.Time
}

func Measure(records []model.Record) Metrics {
	m := Metrics{}
	for _, r := range records {
		m.Seen++
		if r.Status == "draft" {
			m.Rejected++
		} else {
			m.Updated++
		}
		if r.CreatedAt.After(m.LastUpdate) {
			m.LastUpdate = r.CreatedAt
		}
	}
	return m
}
func Filter(records []model.Record, status string) []model.Record {
	out := []model.Record{}
	for _, r := range records {
		if status == "" || r.Status == status {
			out = append(out, r)
		}
	}
	return out
}
func Contains(r model.Record, terms []string) bool {
	hay := strings.ToLower(r.Title + " " + r.Author + " " + strings.Join(r.Tags, " "))
	for _, t := range terms {
		if !strings.Contains(hay, strings.ToLower(t)) {
			return false
		}
	}
	return true
}
func Deduplicate(records []model.Record) []model.Record {
	seen := map[string]bool{}
	out := []model.Record{}
	for _, r := range records {
		if !seen[r.ID] {
			seen[r.ID] = true
			out = append(out, r)
		}
	}
	return out
}
