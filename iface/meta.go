package iface

import (
	_ "embed"
)

//go:embed iface.tmpl
var TmplIface string

const (
	MetaIface   = "+iface.Iface"
	MetaComment = "+iface.Comment"
)

func Name() string {
	return ""
}

//Iface to generate interface for struct
type Iface struct {
	//Name specifies the generated interface name
	Name string
}

//Comment copy comment to generated interface
type Comment struct {
	Value []string
}
