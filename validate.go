package schemata

import (
	"io/ioutil"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

func Validate(pathToCue string, value interface{}) error {

	var ctx = cuecontext.New()
	schemaFileBytes, _ := ioutil.ReadFile(pathToCue)

	ov := ctx.CompileBytes(schemaFileBytes)
	nv := ov.FillPath(cue.ParsePath(""), value)

	return nv.Validate()
}
