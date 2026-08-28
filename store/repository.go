package store

import (
	"bookexchange/model"
	"fmt"
	"time"
)

func (s *Store) SaveProfile(p model.Profile) error { return s.Put("profiles", p.ID, p) }
func (s *Store) LoadProfile(id string) (model.Profile, error) {
	var p model.Profile
	e := s.Get("profiles", id, &p)
	return p, e
}
func (s *Store) SaveEvent(e model.Event) error {
	if e.At.IsZero() {
		e.At = time.Now().UTC()
	}
	return s.Put("events", e.ID, e)
}
func (s *Store) SaveAudit(a model.Audit) error {
	if a.At.IsZero() {
		a.At = time.Now().UTC()
	}
	return s.Put("audits", a.ID, a)
}
func (s *Store) SaveListing(l model.Listing) error {
	return s.Put("listings", fmt.Sprintf("%s:%d", l.ID, l.Version), l)
}
func (s *Store) LoadListing(id string, v int) (model.Listing, error) {
	var l model.Listing
	e := s.Get("listings", fmt.Sprintf("%s:%d", id, v), &l)
	return l, e
}
func (s *Store) SaveNotification(n model.Notification) error { return s.Put("notifications", n.ID, n) }
