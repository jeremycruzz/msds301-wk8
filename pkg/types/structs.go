package types

// WikiResponse is the response from the wikiapi
type WikiResponse struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		Pages map[string]WikiPage `json:"pages"`
	} `json:"query"`
}

// Page contains the needed information
type WikiPage struct {
	PageID  int    `json:"pageid,omitempty"`
	Title   string `json:"title"`
	Extract string `json:"extract,omitempty"` // text content, missing if the page doesn't exist
}

// WikiTableInsert represents the data we want to store in the db
type WikiTableInsert struct {
	PageID  int
	Title   string
	Extract string
	Summary string
	Eli5    string
}

// ChatRequest is a ChatGpt Request
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Message for ChatGpt request
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Chat Response is the Response Object that is returned
type ChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}
