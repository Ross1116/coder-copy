package main

import (
	"fmt"

	"github.com/Ross1116/copy-comment-remover/pkg/monitor"
	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		fmt.Println("Error initialising clipboard:", err)
	}

	fmt.Println("Clipboard monitor started, Press ctrl+C to exit")
	monitor.MonitorClipboard()
}
