package main

import (
	"fmt"
	"os"

	"github.com/Ross1116/coder-copy/pkg/config"
	"github.com/Ross1116/coder-copy/pkg/monitor"
	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		fmt.Println("Error initialising clipboard:", err)
		os.Exit(1)
	}

	cfg := config.GetConfig()
	if cfg != nil {
		fmt.Println("Clipboard monitor started, Press ctrl+C to exit")
		monitor.MonitorClipboard(cfg.Language, cfg.Format)
		return
	}

	processContentFn := func(content, language string, format bool) (string, error) {
		return monitor.ProcessContent(content, language, format)
	}

	p := config.NewProgram(processContentFn)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
