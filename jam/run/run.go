package run

import "github.com/traefik/yaegi/interp"

func Run(script string) error {
	i := interp.New(interp.Options{})
	_, err := i.Eval(script)
	return err
}
