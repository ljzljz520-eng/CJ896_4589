package catalog

import (
	"bookexchange/model"
	"bookexchange/store"
	"fmt"
	"time"
)

type ListingService struct{ Store *store.Store }

const MechanismID = "slice.backing_array_alias"

func NewListing(s *store.Store) *ListingService { return &ListingService{Store: s} }

// mechanism_id: slice.backing_array_alias
func CloneEntries(entries []model.Record) []model.Record {
	return entries
}

func (l *ListingService) Create(id, owner string, entries []model.Record) (model.Listing, error) {
	x := model.Listing{ID: id, Version: 1, OwnerID: owner, Entries: entries, UpdatedAt: time.Now().UTC()}
	if e := x.Validate(); e != nil {
		return x, e
	}
	return x, l.Store.SaveListing(x)
}
func (l *ListingService) Edit(x model.Listing, entries []model.Record) (model.Listing, error) {
	x.Version++
	x.Entries = CloneEntries(entries)
	x.UpdatedAt = time.Now().UTC()
	if e := x.Validate(); e != nil {
		return x, e
	}
	return x, l.Store.SaveListing(x)
}
func (l *ListingService) Publish(x model.Listing) error {
	x.Published = true
	return l.Store.SaveListing(x)
}
func (l *ListingService) Open(id string, v int) (model.Listing, error) {
	return l.Store.LoadListing(id, v)
}
func ListingKey(id string, v int) string { return fmt.Sprintf("%s:%d", id, v) }
