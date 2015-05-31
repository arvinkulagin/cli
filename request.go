package cli

import (
	"strings"
	"errors"
)

type Request struct {
	raw string
	vars map[string]string
}

func NewRequest(raw string, pattern Pattern) (Request, error) {
	request := Request{
		raw: raw,
		vars: make(map[string]string),
	}
	keys := strings.Fields(pattern.Template())
	values := strings.Fields(raw)
	units := strings.Fields(strings.TrimSuffix(strings.TrimPrefix(pattern.RE(), "^"), "$"))
	if len(keys) != len(values) || len(keys) != len(units) {
		return request, errors.New("Wrong command line arguments")
	}
	for i, v := range units {
		if v == "([./0-z]+)" {
			request.vars[strings.TrimSuffix(strings.TrimPrefix(keys[i], "{"), "}")] = values[i]
		}
	}
	return request, nil
}

func (r Request) Vars() map[string]string {
	return r.vars
}

func (r Request) Raw() string {
	return r.raw
}