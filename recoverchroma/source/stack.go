package source

import (
	"fmt"
	"regexp"
	"strings"
)

const sourceRegexp = `\t(.+):([0-9]+)`

var re = regexp.MustCompile(sourceRegexp)

// RenderStack renders in HTML format with links to source code
func RenderStack(stack string) string {
	var b strings.Builder
	lines := strings.Split(stack, "\n")
	for _, line := range lines {
		b.WriteString(renderStackLine(line))
		b.WriteString("\n")
	}
	return b.String()
}

func renderStackLine(line string) string {
	matches := re.FindAllSubmatch([]byte(line), -1)
	if len(matches) < 1 || len(matches[0]) < 2 {
		return line
	}
	p := string(matches[0][1])
	return fmt.Sprintf(`	<a href="/debug/?path=%s">%s</a>`, p, strings.Trim(line, "	"))
}
