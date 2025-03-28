package commentremover

import (
	"bytes"
	"regexp"
	"strings"
)

func CommentRemover(code string, language string) string {
	result := code

	if language == "jsx" {
		result = removeJSXComments(result)
	}

	switch language {
	case "go", "c", "java", "javascript", "js", "jsx":
		return removeCStyleComments(result)
	case "python":
		return removePythonComments(result)
	default:
		return removeCStyleComments(result)
	}
}

func removeJSXComments(code string) string {
	var buffer bytes.Buffer
	inString := false
	stringChar := byte(0)
	i := 0

	for i < len(code) {
		if (code[i] == '"' || code[i] == '\'' || code[i] == '`') && (i == 0 || code[i-1] != '\\') {
			if !inString {
				inString = true
				stringChar = code[i]
			} else if code[i] == stringChar {
				inString = false
			}
			buffer.WriteByte(code[i])
			i++
			continue
		}

		if !inString && i+2 < len(code) && code[i] == '{' && code[i+1] == '/' && code[i+2] == '*' {
			commentEnd := strings.Index(code[i:], "*/}")
			if commentEnd != -1 {
				i += commentEnd + 3
			} else {
				buffer.WriteByte(code[i])
				i++
			}
		} else {
			buffer.WriteByte(code[i])
			i++
		}
	}

	result := buffer.String()

	emptyLinePattern := regexp.MustCompile(`(?m)^\s*\n`)
	result = emptyLinePattern.ReplaceAllString(result, "")

	return result
}

func removeCStyleComments(code string) string {
	var resultLines []string
	lines := strings.Split(code, "\n")

	inMultiLineComment := false

	for i := range lines {
		line := lines[i]
		trimmedLine := strings.TrimSpace(line)

		if len(trimmedLine) == 0 {
			resultLines = append(resultLines, line)
			continue
		}

		processedLine, stillInComment := processCStyleLine(line, inMultiLineComment)
		inMultiLineComment = stillInComment

		if len(strings.TrimSpace(processedLine)) == 0 {
			continue
		}

		resultLines = append(resultLines, processedLine)
	}

	for len(resultLines) > 0 && strings.TrimSpace(resultLines[len(resultLines)-1]) == "" {
		resultLines = resultLines[:len(resultLines)-1]
	}

	return strings.Join(resultLines, "\n")
}

func processCStyleLine(line string, startInComment bool) (string, bool) {
	var result bytes.Buffer
	inString := false
	stringChar := byte(0)
	inComment := startInComment
	i := 0

	for i < len(line) {
		if inComment {
			if i+1 < len(line) && line[i] == '*' && line[i+1] == '/' {
				inComment = false
				i += 2
			} else {
				i++
			}
			continue
		}

		if inString {
			result.WriteByte(line[i])
			if line[i] == '\\' && i+1 < len(line) {
				if i+1 < len(line) {
					result.WriteByte(line[i+1])
				}
				i += 2
			} else if line[i] == stringChar {
				inString = false
				i++
			} else {
				i++
			}
			continue
		}

		if line[i] == '"' || line[i] == '\'' || line[i] == '`' {
			inString = true
			stringChar = line[i]
			result.WriteByte(line[i])
			i++
			continue
		}

		if i+1 < len(line) && line[i] == '/' && line[i+1] == '/' {
			break
		}

		if i+1 < len(line) && line[i] == '/' && line[i+1] == '*' {
			inComment = true
			i += 2
			continue
		}

		result.WriteByte(line[i])
		i++
	}

	return result.String(), inComment
}

func removePythonComments(code string) string {
	var resultLines []string
	lines := strings.Split(code, "\n")

	inTripleQuote := false
	tripleQuoteType := ""

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmedLine := strings.TrimSpace(line)

		if len(trimmedLine) == 0 {
			resultLines = append(resultLines, line)
			continue
		}

		processedLine, stillInTripleQuote, quoteType := processPythonLine(line, inTripleQuote, tripleQuoteType)
		inTripleQuote = stillInTripleQuote
		tripleQuoteType = quoteType

		if len(strings.TrimSpace(processedLine)) == 0 {
			continue
		}

		resultLines = append(resultLines, processedLine)
	}

	for len(resultLines) > 0 && strings.TrimSpace(resultLines[len(resultLines)-1]) == "" {
		resultLines = resultLines[:len(resultLines)-1]
	}

	return strings.Join(resultLines, "\n")
}

func processPythonLine(line string, inTripleQuote bool, tripleQuoteType string) (string, bool, string) {
	if inTripleQuote {
		if strings.Contains(line, tripleQuoteType) {
			idx := strings.Index(line, tripleQuoteType) + 3
			if idx < len(line) {
				return processSinglePythonLine(line[idx:]), false, ""
			}
			return "", false, ""
		}
		return "", true, tripleQuoteType
	}

	if strings.Contains(line, "'''") || strings.Contains(line, "\"\"\"") {
		var quotePos int
		var quoteType string

		if pos := strings.Index(line, "'''"); pos != -1 {
			quoteType = "'''"
			quotePos = pos
		} else {
			quoteType = "\"\"\""
			quotePos = strings.Index(line, "\"\"\"")
		}

		beforeQuote := processSinglePythonLine(line[:quotePos])

		endPos := strings.Index(line[quotePos+3:], quoteType)
		if endPos != -1 {
			endPos += quotePos + 3 + 3

			if endPos < len(line) {
				afterQuote := processSinglePythonLine(line[endPos:])
				return beforeQuote + afterQuote, false, ""
			}
			return beforeQuote, false, ""
		}

		return beforeQuote, true, quoteType
	}

	return processSinglePythonLine(line), false, ""
}

func processSinglePythonLine(line string) string {
	var result bytes.Buffer
	inString := false
	stringChar := byte(0)
	i := 0

	for i < len(line) {
		if !inString && line[i] == '#' {
			break
		}

		if inString {
			result.WriteByte(line[i])
			if line[i] == '\\' && i+1 < len(line) {
				if i+1 < len(line) {
					result.WriteByte(line[i+1])
				}
				i += 2
			} else if line[i] == stringChar {
				inString = false
				i++
			} else {
				i++
			}
			continue
		}

		if line[i] == '"' || line[i] == '\'' {
			inString = true
			stringChar = line[i]
			result.WriteByte(line[i])
			i++
			continue
		}

		result.WriteByte(line[i])
		i++
	}

	return result.String()
}
