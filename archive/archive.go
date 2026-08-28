package archive

import (
	"bookexchange/catalog"
	"bookexchange/model"
	"bookexchange/store"
	"fmt"
	"time"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }
func (a *Service) StoreRecord(id string, c *catalog.Service) error {
	r, e := c.Get(id)
	if e != nil {
		return e
	}
	if e = c.UpdateStatus(id, "archived"); e != nil {
		return e
	}
	return a.Store.SaveAudit(model.Audit{ID: fmt.Sprintf("archive-%s-%d", id, time.Now().UnixNano()), Action: "archive", Entity: "Record", EntityID: id, Detail: r.Title, At: time.Now().UTC()})
}
func (a *Service) Restore(id string, c *catalog.Service) error {
	r, e := c.Get(id)
	if e != nil {
		return e
	}
	if r.Status != "archived" {
		return fmt.Errorf("record is not archived")
	}
	r.Status = "available"
	return a.Store.SaveRecord(r)
}
func (a *Service) Audit(id, action, detail string) error {
	return a.Store.SaveAudit(model.Audit{ID: fmt.Sprintf("audit-%d", time.Now().UnixNano()), Action: action, Entity: "Record", EntityID: id, Detail: detail, At: time.Now().UTC()})
}
