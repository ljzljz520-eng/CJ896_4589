package query

import (
	"bookexchange/catalog"
	"bookexchange/model"
	"sort"
	"strings"
)

type Service struct{ Catalog *catalog.Service }

func New(c *catalog.Service) *Service                  { return &Service{Catalog: c} }
func (q *Service) Find(term string) model.SearchResult { return q.Catalog.Search(term) }
func (q *Service) ByOwner(owner string) []model.Record {
	r := q.Catalog.Search("").Records
	out := []model.Record{}
	for _, x := range r {
		if x.OwnerID == owner {
			out = append(out, x)
		}
	}
	return out
}
func (q *Service) Available() []model.Record {
	r := q.Catalog.Search("").Records
	out := []model.Record{}
	for _, x := range r {
		if x.Status == "available" {
			out = append(out, x)
		}
	}
	return out
}
func (q *Service) SortByEdition(r []model.Record) []model.Record {
	out := append([]model.Record(nil), r...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Edition < out[j].Edition })
	return out
}
func (q *Service) MatchTags(r model.Record, term string) bool {
	term = strings.ToLower(term)
	for _, t := range r.Tags {
		if strings.Contains(strings.ToLower(t), term) {
			return true
		}
	}
	return false
}
