package utils

import (
	"bytes"

	"github.com/yuin/goldmark"
)

func RenderMarkdown(input *string) *string {
	var buf bytes.Buffer

	err := goldmark.Convert([]byte(*input), &buf)
	if err != nil {
		return nil
	}
	result := buf.String()
	return &result
}