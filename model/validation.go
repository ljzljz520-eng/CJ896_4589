package model

import "strings"

func (r Record) Validate() error {
	if strings.TrimSpace(r.ID) == "" {
		return ErrInvalid("record id")
	}
	if strings.TrimSpace(r.Title) == "" {
		return ErrInvalid("title")
	}
	if strings.TrimSpace(r.OwnerID) == "" {
		return ErrInvalid("owner")
	}
	if r.Edition < 0 {
		return ErrInvalid("edition")
	}
	return nil
}
func (p Profile) Validate() error {
	if strings.TrimSpace(p.ID) == "" || strings.TrimSpace(p.Name) == "" {
		return ErrInvalid("profile")
	}
	if !strings.Contains(p.Email, "@") {
		return ErrInvalid("email")
	}
	return nil
}
func (l Listing) Validate() error {
	if l.ID == "" || l.Version < 1 {
		return ErrInvalid("listing")
	}
	if len(l.Entries) == 0 {
		return ErrInvalid("entries")
	}
	return nil
}

type ValidationError struct{ Field string }

func (e ValidationError) Error() string { return "invalid " + e.Field }
func ErrInvalid(f string) error         { return ValidationError{f} }
func NormalizeTags(tags []string) []string {
	out := make([]string, 0, len(tags))
	seen := map[string]bool{}
	for _, t := range tags {
		n := strings.ToLower(strings.TrimSpace(t))
		if n != "" && !seen[n] {
			seen[n] = true
			out = append(out, n)
		}
	}
	return out
}
func StatusAllowed(s string) bool {
	switch s {
	case "draft", "available", "reserved", "exchanged", "archived":
		return true
	default:
		return false
	}
}
