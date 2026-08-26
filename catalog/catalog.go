package catalog

import (
	"bookexchange/model"
	"bookexchange/store"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

type Service struct {
	Store *store.Store
	mu    sync.RWMutex
	cache map[string]model.Record
}

func New(s *store.Store) *Service { return &Service{Store: s, cache: map[string]model.Record{}} }
func (c *Service) Register(r model.Record) error {
	if e := r.Validate(); e != nil {
		return e
	}
	r.Tags = model.NormalizeTags(r.Tags)
	if r.Status == "" {
		r.Status = "draft"
	}
	if r.CreatedAt.IsZero() {
		r.CreatedAt = time.Now().UTC()
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if e := c.Store.SaveRecord(r); e != nil {
		return e
	}
	c.cache[r.ID] = r
	return nil
}
func (c *Service) UpdateStatus(id, status string) error {
	if !model.StatusAllowed(status) {
		return fmt.Errorf("unsupported status")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	r, e := c.Store.LoadRecord(id)
	if e != nil {
		return e
	}
	r.Status = status
	if e = c.Store.SaveRecord(r); e == nil {
		c.cache[id] = r
	}
	return e
}
func (c *Service) Get(id string) (model.Record, error) {
	c.mu.RLock()
	r, ok := c.cache[id]
	c.mu.RUnlock()
	if ok {
		return r, nil
	}
	r, e := c.Store.LoadRecord(id)
	if e == nil {
		c.mu.Lock()
		c.cache[id] = r
		c.mu.Unlock()
	}
	return r, e
}
func (c *Service) Search(q string) model.SearchResult {
	q = strings.ToLower(strings.TrimSpace(q))
	c.mu.RLock()
	all := make([]model.Record, 0, len(c.cache))
	for _, r := range c.cache {
		all = append(all, r)
	}
	c.mu.RUnlock()
	sort.Slice(all, func(i, j int) bool { return all[i].Title < all[j].Title })
	out := make([]model.Record, 0)
	for _, r := range all {
		if q == "" || strings.Contains(strings.ToLower(r.Title), q) || strings.Contains(strings.ToLower(r.Author), q) {
			out = append(out, r)
		}
	}
	return model.SearchResult{Records: out, Total: len(out), Query: q}
}
func (c *Service) Archive(id string) error { return c.UpdateStatus(id, "archived") }
func (c *Service) Snapshot(ids []string) []model.Record {
	out := make([]model.Record, 0, len(ids))
	for _, id := range ids {
		if r, e := c.Get(id); e == nil {
			out = append(out, r)
		}
	}
	return out
}
