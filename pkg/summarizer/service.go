package summarizer

import (
	"github.com/jeremycruzz/msds301-wk8/pkg/types"
)

type Service struct {
	ExistingTitles map[string]bool
	WikiApi        types.WikipediaApi
	Chat           types.ChatGptService
	Db             types.Repository
}

func New(wikiApi types.WikipediaApi, chatgpt types.ChatGptService, db types.Repository) *Service {
	titles := db.GetTitles()
	existingTitles := make(map[string]bool)
	for _, title := range titles {
		existingTitles[title] = true
	}

	return &Service{
		ExistingTitles: existingTitles,
		WikiApi:        wikiApi,
		Chat:           chatgpt,
		Db:             db,
	}
}
