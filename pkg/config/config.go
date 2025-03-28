package config

import (
	"flag"
)

type Config struct {
	Language string
}

func ParseFlags() *Config {

	goLangPtr := flag.Bool("go", true, "Remove Go style comments (default)")
	cLangPtr := flag.Bool("c", false, "Remove C/C++ style comments")
	javaPtr := flag.Bool("java", false, "Remove Java style comments")
	pythonPtr := flag.Bool("python", false, "Remove Python style comments")
	jsPtr := flag.Bool("js", false, "Remove JavaScript style comments")
	flag.Parse()

	language := "go"

	if *cLangPtr {
		language = "c"
	} else if *javaPtr {
		language = "java"
	} else if *pythonPtr {
		language = "python"
	} else if *jsPtr {
		language = "javascript"
	} else if *goLangPtr {
		language = "go"
	}

	return &Config{
		Language: language,
	}
}
