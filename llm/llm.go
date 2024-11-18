package llm

import (
	"bytes"
	"encoding/json"
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

// LLMのAPIサーバーに送信するhttpリクエストの共通処理.
func utilRequest[T any](data T, url string) (*http.Response, error) {
	byteData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json data:%v", err)
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(byteData))
	if err != nil {
		return nil, fmt.Errorf("failed to create new pull request:%v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send pull request:%v", err)
	}
	return resp, nil
}
