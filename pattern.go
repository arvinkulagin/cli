package cli

import (
	"regexp"
)

type Pattern struct {
	re string
	template string
}

func NewPattern(template string) Pattern {
	pattern := Pattern{
		re: "^" + regexp.MustCompile("{([0-z]+)}").ReplaceAllString(template, "([./0-z]+)") + "$",
		template: template,
	}
	return pattern
}

func (p Pattern) RE() string {
	return p.re
}

func (p Pattern) Template() string {
	return p.template
}