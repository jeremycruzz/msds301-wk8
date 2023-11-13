package sqlitedb

import "github.com/jeremycruzz/msds301-wk8/pkg/types"

func (s *service) Get(title string) (*types.WikiTableInsert, error) {
	var wikiInsert types.WikiTableInsert
	err := s.db.QueryRow("SELECT pageid, title, extract, summary, eli5 FROM wiki_pages WHERE title = ?", title).Scan(&wikiInsert.PageID, &wikiInsert.Title, &wikiInsert.Extract, &wikiInsert.Summary, &wikiInsert.Eli5)
	if err != nil {
		return nil, err
	}
	return &wikiInsert, nil
}
