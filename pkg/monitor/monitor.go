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

	if format {
		var lang codeformatter.Language
		switch language {
		case "go":
			lang = codeformatter.Go
		case "cpp", "c++", "c":
			lang = codeformatter.CPP
		case "java":
			lang = codeformatter.Java
		case "javascript", "js":
			lang = codeformatter.JS
		case "typescript", "ts":
			lang = codeformatter.TS
		case "jsx":
			lang = codeformatter.JSX
		case "tsx":
			lang = codeformatter.TSX
		case "python", "py":
			lang = codeformatter.Python
		default:
			lang = codeformatter.Go
		}

		formattedContent, err := codeformatter.FormatCode(strippedContent, lang)
		if err != nil {
			return strippedContent, err
		}
		return formattedContent, nil
	}

	return strippedContent, nil
}
