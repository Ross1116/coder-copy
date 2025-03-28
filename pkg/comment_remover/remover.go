package commentremover

import "regexp"

func CommentRemover(code string, language string) string {
	switch language {
	case "go", "c", "java", "javascript", "js":
		return removeCStyleComments(code)
	case "python":
		return removePythonComments(code)
	default:
		return removeCStyleComments(code)
	}
}

func removeCStyleComments(code string) string {
	multiLinePattern := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	afterMultiLine := multiLinePattern.ReplaceAllString(code, "")

	singleLinePattern := regexp.MustCompile(`(?m)//.*$`)
	afterComments := singleLinePattern.ReplaceAllString(afterMultiLine, "")

	emptyLines := regexp.MustCompile(`(?m)^[ \t]*\r?\n`)
	result := emptyLines.ReplaceAllString(afterComments, "")

	return result
}

func removePythonComments(code string) string {
	singleLinePattern := regexp.MustCompile(`(?m)#.*$`)
	afterSingleLine := singleLinePattern.ReplaceAllString(code, "")

	tripleQuotePattern := regexp.MustCompile(`(?ms)('''.*?'''|""".*?""")`)
	afterComments := tripleQuotePattern.ReplaceAllString(afterSingleLine, "")

	emptyLines := regexp.MustCompile(`(?m)^[ \t]*\r?\n`)
	result := emptyLines.ReplaceAllString(afterComments, "")

	return result
}
