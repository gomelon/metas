package main

import (
	"github.com/gomelon/meta-templates/mwire/testdata/bar"
)

type App struct {
	foo bar.Foo
}

func NewApp(foo bar.Foo) *App {
	return &App{foo: foo}
}
