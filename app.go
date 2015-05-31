package cli

import (
	"regexp"
	"strings"
	"os"
)

var app App = NewApp()

func Add(template string, handler Handler) {
	app.Add(template, handler)
}

func Run() error {
	err := app.Run()
	if err != nil {
		return err
	}
	return nil
}

type Handler func(Request)

type App struct {
	handlers map[Pattern]Handler
}

func NewApp() App {
	app := App{
		handlers: make(map[Pattern]Handler),
	}
	return app
}

func (s *App) Add(template string, handler Handler) {
	pattern := NewPattern(template)
	s.handlers[pattern] = handler
}

func (s *App) Run() error {
	raw := strings.Join(os.Args[1:], " ")
	for pattern, handler := range s.handlers {
		if regexp.MustCompile(pattern.RE()).MatchString(raw) {
			request, err := NewRequest(raw, pattern)
			if err != nil { return err }
			handler(request)
			break
		}
	}
	return nil
}