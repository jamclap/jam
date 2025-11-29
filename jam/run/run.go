package run

import (
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type Runner struct {
	i *interp.Interpreter
}

func NewRunner() *Runner {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	// i.Use(MiniStdExports)
	i.Use(jamExports)
	return &Runner{
		i: i,
	}
}

func (r *Runner) Run(script string) error {
	_, err := r.i.Eval(script)
	return err
}

func (r *Runner) RunPath(path string) error {
	_, err := r.i.EvalPath(path)
	return err
}
