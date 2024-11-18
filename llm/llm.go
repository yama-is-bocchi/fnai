package llm

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type LLM struct {
	baseurl       *url.URL
	model         string
	modelFilePath string
}
type createRequest struct {
	Name      string `json:"name"`
	ModelFile string `json:"modelfile"`
}

type chatRequest struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func New(baseURL *url.URL, modelFilePath string) (LLM, error) {
	if _, err := os.Stat(modelFilePath); err != nil {
		return LLM{}, fmt.Errorf("failed to check model file path:%w", err)
	}
	return LLM{
		baseurl:       baseURL,
		modelFilePath: modelFilePath,
	}, nil
}

func (llm *LLM) CreateModel(modelName string, useModelFile bool) error {
	createReq := createRequest{Name: modelName}
	if useModelFile {
		modelInfo, err := llm.readModelFile()
		if err != nil {
			return fmt.Errorf("failed to read model file:%w", err)
		}
		createReq.ModelFile = modelInfo
	}
	resp, err := utilRequest(createReq,
		llm.baseurl.JoinPath("api", "create").String())
	if err != nil {
		return fmt.Errorf("failed to create model:%w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("an invalid status code was received:%d", resp.StatusCode)
	}
	llm.model = modelName
	return nil
}

func (llm LLM) SendMessage(targetMessage string) (string, error) {
	resp, err := utilRequest(chatRequest{Model: llm.model, Messages: []message{{Role: "user", Content: targetMessage}}},
		llm.baseurl.JoinPath("api", "chat").String())
	if err != nil {
		return "", fmt.Errorf("failed to submit chat request:%w", err)
	}
	respMessage, err := encodeNdJSON(resp)
	if err != nil {
		return "", fmt.Errorf("failed to analysis nd json:%w", err)
	}
	return respMessage, nil
}
