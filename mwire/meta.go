package mwire

import (
	_ "embed"
	"fmt"
	"github.com/gomelon/meta"
	"strconv"
	"text/template"
)

//go:embed wire.tmpl
var TmplWire string

func DefaultPkgGenFactory() meta.PkgGenFactory {
	pkgParser := meta.NewPkgParser()
	return meta.NewTmplPkgGenFactory(TmplWire,
		meta.WithOutputFilename("wire_set"),
		meta.WithPkgParser(pkgParser),
		meta.WithMetaParser(meta.NewParser(pkgParser)),
		meta.WithFuncMapFactory(
			func(gen *meta.TmplPkgGen) template.FuncMap {
				return NewFunctions(gen).FuncMap()
			},
		),
	)
}

const (
	MetaWireProvider = "wire:provider"
)

var (
	MetaNames = []string{MetaWireProvider}
)

//Order It is the injection order,lower values have higher priority(closer it is to the realization).
func Order(m *meta.Meta) int32 {
	orderStr := m.Property("order")
	if len(orderStr) == 0 {
		return 0
	}
	order, err := strconv.ParseInt(orderStr, 10, 32)
	if err != nil {
		panic(fmt.Errorf("wire:provider get order fails: expected a int value,order=%s", orderStr))
	}
	return int32(order)
}
