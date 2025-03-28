package commentremover

import "regexp"

func CommentRemover(code string) string {
	multiLinePattern := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	afterMultiLine := multiLinePattern.ReplaceAllString(code, "")
	singleLinePattern := regexp.MustCompile(`(?m)//.*$`)
	result := singleLinePattern.ReplaceAllString(afterMultiLine, "")

	return result
}
