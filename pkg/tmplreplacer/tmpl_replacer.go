package tmplreplacer

import (
	"regexp"
	"strings"
)

var (
	tmplRe  = regexp.MustCompile(`\${[^{}]*}`)
	spaceRe = regexp.MustCompile(`\s{2,}`)
	punctRe = regexp.MustCompile(`\s+([,.;:!?])`)
)

// TmplReplacer замена шаблонных переменных.
type TmplReplacer struct {
	src string
}

func New(src string) TmplReplacer {
	return TmplReplacer{
		src: src,
	}
}

func (tr TmplReplacer) Replace(varTable map[string]string) string {
	tmpl := tmplRe.ReplaceAllStringFunc(tr.src, func(match string) string {
		varName := match[2 : len(match)-1]
		return varTable[varName]
	})
	tmpl = strings.TrimSpace(tmpl)
	tmpl = spaceRe.ReplaceAllString(tmpl, " ")
	return punctRe.ReplaceAllString(tmpl, "$1")
}
