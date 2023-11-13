package sqlitedb

import (
	"strings"

	"github.com/jeremycruzz/msds301-wk8/pkg/types"
)

func (s *service) Insert(wikiInsert *types.WikiTableInsert) error {
	_, err := s.db.Exec("INSERT INTO wiki_pages (pageid, title, extract, summary, eli5) VALUES (?, ?, ?, ?, ?)", wikiInsert.PageID, strings.ToLower(wikiInsert.Title), wikiInsert.Extract, wikiInsert.Summary, wikiInsert.Eli5)
	return err
}
