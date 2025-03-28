package codeformatter

import (
	"fmt"
	"go/format"
)

func FormatCode(code string) (string, error) {
	formattedBytes, err := format.Source([]byte(code))
	if err != nil {
		return code, fmt.Errorf("formatting error: %v", err)
	}

	return string(formattedBytes), nil
}
