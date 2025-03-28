package monitor

import (
	"context"
	"fmt"

	codeformatter "github.com/Ross1116/copy-comment-remover/pkg/code_formatter"
	commentremover "github.com/Ross1116/copy-comment-remover/pkg/comment_remover"
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
			var finalWrite string
			var err error

			strippedContent := commentremover.CommentRemover(currContent, language)
			if format && language == "go" {
				finalWrite, err = codeformatter.FormatCode(strippedContent)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				finalWrite = strippedContent
			}
			clipboard.Write(clipboard.FmtText, []byte(finalWrite))
			prevContent = currContent
		}
	}

	return prevContent
}
