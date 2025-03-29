package monitor

import (
	"context"
	"fmt"

	codeformatter "github.com/Ross1116/coder-copy/pkg/code_formatter"
	commentremover "github.com/Ross1116/coder-copy/pkg/comment_remover"
	"golang.design/x/clipboard"
)

func MonitorClipboard(language string, format bool) string {
	ctx := context.Background()
	copied := clipboard.Watch(ctx, clipboard.FmtText)
	var prevContent string

	for content := range copied {
		currContent := string(content)

		if currContent != prevContent {
			fmt.Println(currContent)
			processedContent, _ := ProcessContent(currContent, language, format)
			clipboard.Write(clipboard.FmtText, []byte(processedContent))
			prevContent = currContent
		}
	}

	return prevContent
}

func ProcessContent(content, language string, format bool) (string, error) {
	strippedContent := commentremover.CommentRemover(content, language)

	if format && language == "go" {
		formattedContent, err := codeformatter.FormatCode(strippedContent)
		if err != nil {
			return strippedContent, err
		}
		return formattedContent, nil
	}

	return strippedContent, nil
}
