package reports

import (
	"bookexchange/model"
	"strings"
)

type Policy struct {
	MaxTags         int
	RequireAuthor   bool
	AllowedStatuses map[string]bool
}

func DefaultPolicy() Policy {
	return Policy{MaxTags: 8, RequireAuthor: true, AllowedStatuses: map[string]bool{"draft": true, "available": true, "reserved": true, "exchanged": true, "archived": true}}
}
func (p Policy) Check(r model.Record) []string {
	issues := []string{}
	if p.RequireAuthor && strings.TrimSpace(r.Author) == "" {
		issues = append(issues, "author required")
	}
	if p.MaxTags > 0 && len(r.Tags) > p.MaxTags {
		issues = append(issues, "too many tags")
	}
	if !p.AllowedStatuses[r.Status] {
		issues = append(issues, "status forbidden")
	}
	return issues
}
func (p Policy) Accept(r model.Record) bool { return len(p.Check(r)) == 0 }
func MergePolicies(a, b Policy) Policy {
	out := a
	if b.MaxTags > 0 && (out.MaxTags == 0 || b.MaxTags < out.MaxTags) {
		out.MaxTags = b.MaxTags
	}
	out.RequireAuthor = a.RequireAuthor || b.RequireAuthor
	if out.AllowedStatuses == nil {
		out.AllowedStatuses = map[string]bool{}
	}
	for k, v := range b.AllowedStatuses {
		out.AllowedStatuses[k] = v
	}
	return out
}
