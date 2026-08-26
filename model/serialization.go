package model

import (
	"encoding/json"
	"time"
)

func EncodeRecord(r Record) ([]byte, error) { return json.Marshal(r) }
func DecodeRecord(b []byte) (Record, error) { var r Record; e := json.Unmarshal(b, &r); return r, e }
func NewEvent(id, record, kind, actor, msg string) Event {
	return Event{ID: id, RecordID: record, Kind: kind, ActorID: actor, Message: msg, At: time.Now().UTC()}
}
func NewAudit(id, action, entity, eid, detail string) Audit {
	return Audit{ID: id, Action: action, Entity: entity, EntityID: eid, Detail: detail, At: time.Now().UTC()}
}
