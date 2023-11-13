package sqlitedb

import (
	"log"

	"github.com/jeremycruzz/msds301-wk8/pkg/types"
)

func (r *service) Update(data *types.WikiTableInsert) error {
	query := `UPDATE wiki_pages 
              SET summary = ?, eli5 = ? 
              WHERE title = ?`

	_, err := r.db.Exec(query, data.Summary, data.Eli5, data.Title)
	if err != nil {
		log.Printf("Error updating record: %v", err)
		return err
	}

	return nil
}
