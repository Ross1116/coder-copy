package config

import (
	"flag"
	"os"
)

type Config struct {
	Language string
	Format   bool
}

func GetConfig() *Config {
	if len(os.Args) > 1 {
		return parseFlags()
	}

	return nil
}

func parseFlags() *Config {
	goPtr := flag.Bool("go", false, "Remove Go style comments")
	cLangPtr := flag.Bool("c", false, "Remove C/C++ style comments")
	javaPtr := flag.Bool("java", false, "Remove Java style comments")
	pythonPtr := flag.Bool("python", false, "Remove Python style comments")
	jsPtr := flag.Bool("js", false, "Remove JavaScript style comments")
	jsxPtr := flag.Bool("jsx", false, "Remove JSX style comments")
	formatPtr := flag.Bool("format", false, "Format the copied code automatically")
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
	} else if *jsxPtr {
		language = "jsx"
	} else if *goPtr {
		language = "go"
	}

	return &Config{
		Language: language,
		Format:   *formatPtr,
	}
}
