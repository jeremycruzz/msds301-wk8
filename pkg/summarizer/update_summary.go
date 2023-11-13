package summarizer

import (
	"errors"

	"github.com/jeremycruzz/msds301-wk8/pkg/types"
)

func (s *Service) UpdateSummary(topic string) (summary string, eli5 string, err error) {
	if !s.ExistingTitles[topic] {
		return "", "", errors.New("topic does not exist")
	}

	// get from db
	existingData, err := s.Db.Get(topic)
	if err != nil {
		return "", "", err
	}

	newSummary, newEli5, err := s.Chat.Ask(existingData.Extract)
	if err != nil {
		return "", "", err
	}

	// store in DB
	err = s.Db.Update(&types.WikiTableInsert{
		PageID:  existingData.PageID,
		Title:   existingData.Title,
		Extract: existingData.Extract,
		Summary: newSummary,
		Eli5:    newEli5,
	})
	if err != nil {
		return "", "", err
	}

	return newSummary, newEli5, nil
}
