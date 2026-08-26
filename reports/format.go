package reports

import (
	"bookexchange/model"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func WriteCSV(w io.Writer, records []model.Record) error {
	c := csv.NewWriter(w)
	if e := c.Write([]string{"id", "title", "author", "owner", "status", "edition", "tags"}); e != nil {
		return e
	}
	for _, r := range records {
		if e := c.Write([]string{r.ID, r.Title, r.Author, r.OwnerID, r.Status, strconv.Itoa(r.Edition), strings.Join(r.Tags, "|")}); e != nil {
			return e
		}
	}
	c.Flush()
	return c.Error()
}
func ReadCSV(r io.Reader) ([]model.Record, error) {
	c := csv.NewReader(r)
	if _, e := c.Read(); e != nil {
		return nil, e
	}
	out := []model.Record{}
	for {
		row, e := c.Read()
		if e == io.EOF {
			break
		}
		if e != nil {
			return nil, e
		}
		if len(row) < 7 {
			return nil, fmt.Errorf("invalid row")
		}
		edition, e := strconv.Atoi(row[5])
		if e != nil {
			return nil, e
		}
		out = append(out, model.Record{ID: row[0], Title: row[1], Author: row[2], OwnerID: row[3], Status: row[4], Edition: edition, Tags: strings.Split(row[6], "|")})
	}
	return out, nil
}
func Quote(s string) string { return `"` + strings.ReplaceAll(s, `"`, `""`) + `"` }
func SafeTitle(s string) string {
	if strings.TrimSpace(s) == "" {
		return "未命名"
	}
	return strings.TrimSpace(s)
}
