package monitor

import (
	"context"

	commentremover "github.com/Ross1116/copy-comment-remover/pkg/comment_remover"
	"golang.design/x/clipboard"
)

func MonitorClipboard(language string) string {
	ctx := context.Background()
	copied := clipboard.Watch(ctx, clipboard.FmtText)
	var prevContent string

	for content := range copied {
		currContent := string(content)

		if currContent != prevContent {
			// fmt.Println(currContent)
			strippedContent := commentremover.CommentRemover(currContent, language)
			clipboard.Write(clipboard.FmtText, []byte(strippedContent))
			prevContent = currContent
		}
	}

	return prevContent
}
