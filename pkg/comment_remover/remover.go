package commentremover

import "regexp"

func CommentRemover(code string, language string) string {
	switch language {
	case "go", "c", "java", "javascript", "js":
		return removeCStyleComments(code, language)
	case "python":
		return removePythonComments(code)
	default:
		return removeCStyleComments(code, language)
	}
}

func removeCStyleComments(code string, language string) string {
	var afterJSXComments string

	if language == "jsx" {
		jsxCommentPattern := regexp.MustCompile(`\{/\*[\s\S]*?\*/\}`)
		afterJSXComments = jsxCommentPattern.ReplaceAllString(code, "")
	} else {
		afterJSXComments = code
	}

	multiLinePattern := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	afterMultiLine := multiLinePattern.ReplaceAllString(afterJSXComments, "")

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
	result := emptyLines.ReplaceAllString(afterComments, "\n")

	return result
}
