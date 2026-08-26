package api

import (
	"bookexchange/model"
	"bookexchange/workflow"
	"encoding/json"
	"net/http"
	"strings"
)

type Server struct{ Workflow *workflow.Service }

func New(w *workflow.Service) *Server { return &Server{Workflow: w} }
func (s *Server) Handler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/records", s.records)
	m.HandleFunc("/health", func(x http.ResponseWriter, _ *http.Request) { x.WriteHeader(http.StatusOK); x.Write([]byte("ok")) })
	return m
}
func (s *Server) records(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var x model.Record
		if json.NewDecoder(r.Body).Decode(&x) != nil {
			s.writeError(w, http.StatusBadRequest, "invalid json")
			return
		}
		if e := s.Workflow.Receive(x); e != nil {
			s.writeError(w, http.StatusBadRequest, e.Error())
			return
		}
		s.writeJSON(w, http.StatusCreated, x)
		return
	}
	q := s.Workflow.QueryRecords(r.URL.Query().Get("q"))
	s.writeJSON(w, http.StatusOK, q.Records)
}
func (s *Server) writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
func (s *Server) writeError(w http.ResponseWriter, code int, msg string) {
	s.writeJSON(w, code, map[string]string{"error": strings.TrimSpace(msg)})
}
