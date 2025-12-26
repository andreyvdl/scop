package src

import (
	"strings"
)

func ClearLineContent(line string) string {
	if line == "" || line[0] == '#' {
		return ""
	}

	line, _, _ = strings.Cut(line, "#")
	line = strings.TrimSpace(line)
	return line
}
