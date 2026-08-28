package workflow

import (
	"bookexchange/archive"
	"bookexchange/catalog"
	"bookexchange/model"
	"bookexchange/query"
	"bookexchange/store"
	"fmt"
	"time"
)

type Service struct {
	Catalog  *catalog.Service
	Listings *catalog.ListingService
	Store    *store.Store
	Archive  *archive.Service
	Query    *query.Service
}

func New(s *store.Store) *Service {
	c := catalog.New(s)
	return &Service{Catalog: c, Listings: catalog.NewListing(s), Store: s, Archive: archive.New(s), Query: query.New(c)}
}
func (w *Service) Receive(r model.Record) error {
	if r.Status == "" {
		r.Status = "draft"
	}
	return w.Catalog.Register(r)
}
func (w *Service) Process(id string, actor string) error {
	if e := w.Catalog.UpdateStatus(id, "available"); e != nil {
		return e
	}
	return w.Store.SaveEvent(model.Event{ID: fmt.Sprintf("process-%s", id), RecordID: id, Kind: "processed", ActorID: actor, At: time.Now().UTC()})
}
func (w *Service) Exchange(id, actor string) error {
	if e := w.Catalog.UpdateStatus(id, "exchanged"); e != nil {
		return e
	}
	return w.Store.SaveEvent(model.Event{ID: fmt.Sprintf("exchange-%s", id), RecordID: id, Kind: "exchanged", ActorID: actor, At: time.Now().UTC()})
}
func (w *Service) ArchiveRecord(id string) error            { return w.Archive.StoreRecord(id, w.Catalog) }
func (w *Service) QueryRecords(q string) model.SearchResult { return w.Query.Find(q) }
func (w *Service) Notify(profile, subject, body string) error {
	return w.Store.SaveNotification(model.Notification{ID: fmt.Sprintf("%s-%d", profile, time.Now().UnixNano()), ProfileID: profile, Subject: subject, Body: body, SentAt: time.Now().UTC()})
}
