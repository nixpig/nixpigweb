package utils

import (
	"bytes"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

func MdToHtml(md []byte) (string, error) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("rrt"),
				highlighting.WithFormatOptions(chromahtml.WithLineNumbers(true)))))

	var buf bytes.Buffer
	if err := markdown.Convert(md, &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}
