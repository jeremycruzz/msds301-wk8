package sqlitedb

import "log"

func (s *service) GetTitles() []string {
	var titles []string

	rows, err := s.db.Query("SELECT title FROM wiki_pages")
	if err != nil {
		log.Fatalf("Error querying titles: %v", err)
		return titles
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			log.Fatalf("Error scanning title: %v", err)
			return titles
		}
		titles = append(titles, title)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error during rows iteration: %v", err)
		return titles
	}

	return titles
}
