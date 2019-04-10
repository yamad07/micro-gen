package generator

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func Entity(name string) error {
	f := NewFile("models")
	f.Type().Id(strcase.ToCamel(name)).Struct()
	err := f.Save("entities/models/" + strcase.ToSnake(name) + ".go")
	if err != nil {
		return err
	}
	fmt.Printf("\x1b[33m%s\x1b[0m entities/models/"+strcase.ToSnake(name)+".go\n", "create")

	f = NewFile("repositories")
	f.Type().Id(strcase.ToCamel(name) + "Repository").Interface()
	f.Type().Id(strcase.ToCamel(name) + "RepositoryImpl").Struct()
	f.Func().Id("New" + strcase.ToCamel(name) + "RepositoryImpl").Params().Id(strcase.ToCamel(name) + "RepositoryImpl").Block()
	err = f.Save("entities/repositories/" + strcase.ToSnake(name) + ".go")
	if err != nil {
		return err
	}
	fmt.Printf("\x1b[33m%s\x1b[0m entities/repositories/"+strcase.ToSnake(name)+".go\n", "create")
	return nil
}
