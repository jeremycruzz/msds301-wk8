package summarizer

import "github.com/jeremycruzz/msds301-wk8/pkg/types"

func (s *Service) Summarize(topic string) (summary string, eli5 string, err error) {
	if s.ExistingTitles[topic] {
		data, err := s.Db.Get(topic)
		if err != nil {
			return "", "", err
		}
		return data.Summary, data.Eli5, nil
	}

	// get from wikiapi
	wikiPage, err := s.WikiApi.Query(topic)
	if err != nil {
		return "", "", err
	}

	// ask ChatGPT for summary and ELI5
	summary, eli5, err = s.Chat.Ask(wikiPage.Extract)
	if err != nil {
		return "", "", err
	}

	// store in DB
	err = s.Db.Insert(&types.WikiTableInsert{
		PageID:  wikiPage.PageID,
		Title:   wikiPage.Title,
		Extract: wikiPage.Extract,
		Summary: summary,
		Eli5:    eli5,
	})
	if err != nil {
		return "", "", err
	}

	// add to existing titles
	s.ExistingTitles[topic] = true

	return summary, eli5, nil
}
