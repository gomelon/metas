// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gomelon/meta-templates/mwire/testdata/bar"
)

// Injectors from wire.go:

func initApp(greeting bar.Greeting, bye bar.Bye) (*App, error) {
	defaultFoo := bar.NewDefaultFoo()
	fooWithGreet := bar.NewFooWithGreet(defaultFoo, greeting)
	fooWithBye := bar.NewFooWithBye(fooWithGreet, bye)
	app := NewApp(fooWithBye)
	return app, nil
}
