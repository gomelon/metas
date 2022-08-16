package iface

import (
	"fmt"
	"github.com/gomelon/meta"
	"os"
	"testing"
)

func TestTemplateGen(t *testing.T) {
	workdir, _ := os.Getwd()
	path := workdir + "/testdata"
	generator, err := meta.NewTmplPkgGen(path, TmplIface)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = generator.Generate()
	if err != nil {
		fmt.Println(err.Error())
	}
}
