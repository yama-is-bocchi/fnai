package llm

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type chatResponse struct {
	Model      string    `json:"model"`
	Created_at time.Time `json:"created_at"`
	Message    message   `json:"message"`
	Done       bool      `json:"done"`
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

func encodeNdJSON(aResponse *http.Response) (string, error) {
	var tContents []string
	tScanner := bufio.NewScanner(aResponse.Body)
	for tScanner.Scan() {
		var tResponseData chatResponse
		if tError := json.Unmarshal(tScanner.Bytes(), &tResponseData); tError != nil {
			return "", fmt.Errorf("failed to unmarshal response body:%v", tError)
		}
		tContents = append(tContents, tResponseData.Message.Content)
	}
	if tError := tScanner.Err(); tError != nil {
		return "", fmt.Errorf("failed to read response:%v", tError)
	}
	return strings.Join(tContents, ""), nil
}
