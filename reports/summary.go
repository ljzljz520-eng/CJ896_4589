package reports

import (
	"bookexchange/model"
	"sort"
	"time"
)

type Summary struct {
	GeneratedAt                           time.Time
	Total, Available, Archived, Exchanged int
	ByOwner                               map[string]int
}

func BuildSummary(records []model.Record) Summary {
	s := Summary{GeneratedAt: time.Now().UTC(), ByOwner: map[string]int{}}
	for _, r := range records {
		s.Total++
		s.ByOwner[r.OwnerID]++
		switch r.Status {
		case "available":
			s.Available++
		case "archived":
			s.Archived++
		case "exchanged":
			s.Exchanged++
		}
	}
	return s
}
func TopOwners(s Summary, n int) []string {
	keys := make([]string, 0, len(s.ByOwner))
	for k := range s.ByOwner {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if s.ByOwner[keys[i]] == s.ByOwner[keys[j]] {
			return keys[i] < keys[j]
		}
		return s.ByOwner[keys[i]] > s.ByOwner[keys[j]]
	})
	if n > len(keys) {
		n = len(keys)
	}
	return keys[:n]
}
func StatusLabel(status string) string {
	switch status {
	case "available":
		return "可交换"
	case "reserved":
		return "已预留"
	case "exchanged":
		return "已交换"
	case "archived":
		return "已归档"
	default:
		return "草稿"
	}
}
func CloneRecords(in []model.Record) []model.Record {
	out := make([]model.Record, len(in))
	for i, r := range in {
		out[i] = r
		out[i].Tags = append([]string(nil), r.Tags...)
	}
	return out
}
