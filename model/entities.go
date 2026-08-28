package model

import "time"

type Record struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	OwnerID   string    `json:"owner_id"`
	Status    string    `json:"status"`
	Edition   int       `json:"edition"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
}
type Profile struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Reputation int    `json:"reputation"`
	Active     bool   `json:"active"`
}
type Event struct {
	ID       string    `json:"id"`
	RecordID string    `json:"record_id"`
	Kind     string    `json:"kind"`
	ActorID  string    `json:"actor_id"`
	Message  string    `json:"message"`
	At       time.Time `json:"at"`
}
type Audit struct {
	ID       string    `json:"id"`
	Action   string    `json:"action"`
	Entity   string    `json:"entity"`
	EntityID string    `json:"entity_id"`
	Detail   string    `json:"detail"`
	At       time.Time `json:"at"`
}
type Listing struct {
	ID        string    `json:"id"`
	Version   int       `json:"version"`
	OwnerID   string    `json:"owner_id"`
	Entries   []Record  `json:"entries"`
	Published bool      `json:"published"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Notification struct {
	ID        string    `json:"id"`
	ProfileID string    `json:"profile_id"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	Read      bool      `json:"read"`
	SentAt    time.Time `json:"sent_at"`
}
type SearchResult struct {
	Records []Record
	Total   int
	Query   string
}
