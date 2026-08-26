package workflow

import (
	"bookexchange/model"
	"fmt"
	"time"
)

type Step struct {
	Name  string
	At    time.Time
	OK    bool
	Error string
}
type Chain struct {
	ID       string
	Steps    []Step
	Complete bool
}

func StartChain(id string) Chain { return Chain{ID: id, Steps: []Step{}} }
func (c *Chain) Add(name string, ok bool, err error) {
	s := Step{Name: name, At: time.Now().UTC(), OK: ok}
	if err != nil {
		s.Error = err.Error()
	}
	c.Steps = append(c.Steps, s)
	c.Complete = ok && c.CompleteOrFirst()
}
func (c Chain) CompleteOrFirst() bool {
	if len(c.Steps) == 0 {
		return false
	}
	for _, s := range c.Steps {
		if !s.OK {
			return false
		}
	}
	return true
}
func (c Chain) Failed() bool {
	for _, s := range c.Steps {
		if !s.OK {
			return true
		}
	}
	return false
}
func (c Chain) Summary() string {
	if c.Complete {
		return fmt.Sprintf("%s complete", c.ID)
	}
	return fmt.Sprintf("%s pending", c.ID)
}
func (w *Service) RunIntake(r model.Record) (Chain, error) {
	c := StartChain("intake-" + r.ID)
	e := w.Receive(r)
	c.Add("receive", e == nil, e)
	if e != nil {
		return c, e
	}
	e = w.Process(r.ID, r.OwnerID)
	c.Add("process", e == nil, e)
	return c, e
}
