package main

import (
	"fmt"
	"os"

	"github.com/Ross1116/copy-comment-remover/pkg/config"
	"github.com/Ross1116/copy-comment-remover/pkg/monitor"
	"golang.design/x/clipboard"
)

func main() {
	config := config.GetConfig()

	err := clipboard.Init()
	if err != nil {
		fmt.Println("Error initialising clipboard:", err)
		os.Exit(1)
	}

	fmt.Println("Clipboard monitor started, Press ctrl+C to exit")
	monitor.MonitorClipboard(config.Language, config.Format)
}
