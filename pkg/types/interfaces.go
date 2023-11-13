//go:generate mockgen -source=interfaces.go -destination=../mocks/mocks.go -package=mocks

package types

type WikipediaApi interface {
	Query(query string) (*WikiPage, error)
}

type Repository interface {
	GetTitles() []string
	Get(title string) (*WikiTableInsert, error)
	Insert(*WikiTableInsert) error
	Update(data *WikiTableInsert) error
}

type ChatGptService interface {
	Ask(prompt string) (summary string, eli5 string, err error)
}

type SummarizerService interface {
	Summarize(topic string) (string, string, error)
}
