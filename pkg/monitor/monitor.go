package monitor

import (
	"context"
	"fmt"

	"golang.design/x/clipboard"
)

func MonitorClipboard() {
	ctx := context.Background()
	copied := clipboard.Watch(ctx, clipboard.FmtText)

	var prevContent string

	for content := range copied {
		currContent := string(content)

		if currContent != prevContent {
			fmt.Println(currContent)
			prevContent = currContent
		}
	}
}
