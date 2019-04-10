package generator

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func Controller(name string, action string) error {
	f := NewFile("controllers")
	f.Type().Id(strcase.ToCamel(name) + "Controller").Interface(
		Id(strcase.ToCamel(action)).Params(
			Id("c").Id("context.Context"),
		).Params(
			Id("err").Id("error"),
		),
	)
	err := f.Save("controllers/" + strcase.ToSnake(name) + ".go")
	fmt.Printf("\x1b[33m%s\x1b[0m controllers/"+strcase.ToSnake(name)+".go\n", "create")
	return err
}
