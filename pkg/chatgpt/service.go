package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jeremycruzz/msds301-wk8/pkg/types"
)

const SUMMARY_PROMPT = "Can you summarize this wikipedia article with bullet points?. "
const ELI5_PROMPT = "Can you explain this summary about a wikipedia article as if I were a five year old?. "
const API_ENDPOINT = "https://api.openai.com/v1/chat/completions"
const GPT4_PREVIEW = "gpt-4-1106-preview"

type service struct {
	apiKey string
	model  string
	url    string
}

func New(apiKey string) *service {
	return &service{
		apiKey: apiKey,
		model:  GPT4_PREVIEW,
		url:    API_ENDPOINT,
	}
}

func (s *service) Ask(text string) (summary, eli5 string, err error) {
	summary, err = s.askSummary(text)
	if err != nil {
		return "", "", fmt.Errorf("error getting summary: %s", err)
	}

	// the summary is used instead of the text here to avoid using up my tokens since
	// the wiki text is super long
	eli5, err = s.askEli5(summary)
	if err != nil {
		return "", "", fmt.Errorf("error getting eli5: %s", err)
	}

	return summary, eli5, nil

}

func (s *service) askSummary(prompt string) (string, error) {
	var builder strings.Builder
	builder.WriteString(SUMMARY_PROMPT)
	builder.WriteString(prompt)

	resp, err := s.AskCustom(builder.String())
	if err != nil {
		return "", err
	}

	return resp, nil
}

func (s *service) askEli5(prompt string) (string, error) {
	var builder strings.Builder
	builder.WriteString(ELI5_PROMPT)
	builder.WriteString(prompt)

	resp, err := s.AskCustom(builder.String())
	if err != nil {
		return "", err
	}

	return resp, nil
}

func (s *service) AskCustom(prompt string) (string, error) {
	// construct request
	chatReq := types.ChatRequest{
		Model: s.model,
		Messages: []types.Message{
			{
				Role:    "user", //always user
				Content: prompt,
			},
		},
	}

	reqBody, err := json.Marshal(chatReq)
	if err != nil {
		return "", err
	}

	// create request
	req, err := http.NewRequest("POST", s.url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	// exectute
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatResp types.ChatResponse
	err = json.Unmarshal(body, &chatResp)
	if err != nil {
		return "", err
	}

	// get response
	if len(chatResp.Choices) > 0 && len(chatResp.Choices[0].Message.Content) > 0 {
		return chatResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from ChatGPT")
}
