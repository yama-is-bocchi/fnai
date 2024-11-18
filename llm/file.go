package llm

import (
	"fmt"
	"os"
)

func (llm LLM) readModelFile() (string, error) {
	data, err := os.ReadFile(llm.modelFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to open model file:%v", err)
	}
	return string(data), nil
}
