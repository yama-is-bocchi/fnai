FROM ollama/ollama:latest

RUN apt-get update && apt-get install -y curl

RUN mkdir -p /model

RUN curl -L -o /model/ggml-model-Q4_K_M.gguf https://huggingface.co/lightblue/suzume-llama-3-8B-japanese-gguf/resolve/main/ggml-model-Q4_K_M.gguf
