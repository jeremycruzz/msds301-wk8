package wikiapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jeremycruzz/msds301-wk8/pkg/types"
)

const QUERY_FORMAT_STRING = "https://en.wikipedia.org/w/api.php?action=query&origin=*&prop=extracts&explaintext&titles=%s&format=json"

type service struct {
}

func New() *service {
	return &service{}
}

func (*service) Query(topic string) (*types.WikiPage, error) {
	url := fmt.Sprintf(QUERY_FORMAT_STRING, topic)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var wikiResponse types.WikiResponse
	err = json.Unmarshal(body, &wikiResponse)
	if err != nil {
		return nil, err
	}

	for _, page := range wikiResponse.Query.Pages {
		// return first page for now
		return &page, nil
	}

	return nil, fmt.Errorf("no pages found for topic: %s", topic)
}
