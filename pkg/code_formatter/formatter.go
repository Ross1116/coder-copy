package codeformatter

import (
	"fmt"
	"go/format"
	"os/exec"
	"strings"
)

type Language string

const (
	Go     Language = "go"
	CPP    Language = "cpp"
	Java   Language = "java"
	JS     Language = "javascript"
	TS     Language = "typescript"
	JSX    Language = "jsx"
	TSX    Language = "tsx"
	Python Language = "python"
)

func FormatCode(code string, lang Language) (string, error) {
	switch lang {
	case Go:
		return formatGo(code)
	case CPP:
		return formatCPP(code)
	case Java:
		return formatJava(code)
	case JS, TS:
		return formatJavaScript(code, "babel")
	case JSX:
		return formatJavaScript(code, "babel")
	case TSX:
		return formatJavaScript(code, "typescript")
	case Python:
		return formatPython(code)
	default:
		return formatGo(code)
	}
}

func formatGo(code string) (string, error) {
	formattedBytes, err := format.Source([]byte(code))
	if err != nil {
		return code, fmt.Errorf("go formatting error: %v", err)
	}
	return string(formattedBytes), nil
}
func formatCPP(code string) (string, error) {
	return formatWithExternalTool(code, "clang-format", []string{}, "C/C++",
		"Install clang-format: https://clang.llvm.org/docs/ClangFormat.html")
}

func formatJava(code string) (string, error) {
	if _, err := exec.LookPath("google-java-format"); err == nil {
		return formatWithExternalTool(code, "google-java-format", []string{"-"}, "Java",
			"Install google-java-format: https://github.com/google/google-java-format")
	}

	if _, err := exec.LookPath("java"); err == nil {
		return formatWithExternalTool(code, "java", []string{"-jar", "/usr/local/lib/google-java-format.jar", "-"}, "Java",
			"Install google-java-format: https://github.com/google/google-java-format")
	}

	return code, fmt.Errorf("java formatter not found")
}

func formatJavaScript(code string, parser string) (string, error) {
	return formatWithExternalTool(code, "prettier", []string{"--stdin", "--parser", parser},
		fmt.Sprintf("JavaScript (%s)", parser),
		"Install prettier: npm install -g prettier")
}

func formatPython(code string) (string, error) {
	if _, err := exec.LookPath("black"); err == nil {
		return formatWithExternalTool(code, "black", []string{"-", "-q"}, "Python",
			"Install black: pip install black")
	}

	if _, err := exec.LookPath("python"); err == nil {
		return formatWithExternalTool(code, "python", []string{"-m", "black", "-", "-q"}, "Python",
			"Install black: pip install black")
	}

	if _, err := exec.LookPath("python3"); err == nil {
		return formatWithExternalTool(code, "python3", []string{"-m", "black", "-", "-q"}, "Python",
			"Install black: pip install black")
	}

	return code, fmt.Errorf("python formatter not found")
}

func formatWithExternalTool(code, command string, args []string, language, installInstructions string) (string, error) {
	_, err := exec.LookPath(command)
	if err != nil {
		return code, fmt.Errorf("%s formatter not found. %s", language, installInstructions)
	}

	cmd := exec.Command(command, args...)
	cmd.Stdin = strings.NewReader(code)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return code, fmt.Errorf("%s formatting error: %v\n%s", language, err, output)
	}
	return string(output), nil
}
