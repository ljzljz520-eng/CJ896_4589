package api

import (
	"bookexchange/store"
	"bookexchange/workflow"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestHealth(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	New(workflow.New(s)).Handler().ServeHTTP(r, req)
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
