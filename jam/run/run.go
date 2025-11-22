package run

import (
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func Run(script string) error {
	i := interp.New(interp.Options{})
	_, err := i.Eval(script)
	return err
}

func RunPath(path string) error {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	_, err := i.EvalPath(path)
	return err
}
