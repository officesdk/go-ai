package aimodel

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type ChatCompletionStream struct {
	*StreamReader
}

type StreamReader struct {
	isFinished bool
	Reader     *bufio.Reader
}

func (stream *StreamReader) Recv() (response ChatCompletionStreamResponse, err error) {
	rawLine, err := stream.RecvRaw()
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(rawLine), &response)
	if err != nil {
		return
	}
	return response, nil
}

func (stream *StreamReader) RecvRaw() (string, error) {
	if stream.isFinished {
		return "", io.EOF
	}

	return stream.processLines()
}

//nolint:gocognit
func (stream *StreamReader) processLines() (string, error) {
	for {
		line, err := stream.Reader.ReadString('\n') // Read until newline
		if err != nil {
			if err == io.EOF {
				return "", io.EOF
			}
			return "", fmt.Errorf("error reading stream: %w", err)
		}

		line = strings.TrimSpace(line)
		if line == "data: [DONE]" {
			stream.isFinished = true
			return "", io.EOF // End of stream
		}
		if len(line) > 6 && line[:6] == "data: " {
			trimmed := line[6:] // Trim the "data: " prefix
			return trimmed, nil
		}
	}
}
