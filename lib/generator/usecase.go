package generator

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func Usecase(name string, action string) error {
	f := NewFile("usecases")
	f.Type().Id(strcase.ToCamel(name) + "Usecase").Interface(
		Id(strcase.ToCamel(action)).Params(
			Id("ipt").Id("inputs." + strcase.ToCamel(action)),
		).Params(
			Id("opt").Id("outputs." + strcase.ToCamel(action)),
		),
	)

	iname := strcase.ToSnake(name) + "UsecaseInteractor"
	f.Type().Id(iname).Struct()
	f.Func().Id("New" + strcase.ToCamel(name) + "UsecaseInteractor").Params().Id(strcase.ToCamel(name) + "UsecaseInteractor").Block()

	f.Func().Params(
		Id("i").Id(iname),
	).Id(strcase.ToCamel(action)).Params(
		Id("ipt").Id("inputs." + strcase.ToCamel(action)),
	).Params(
		Id("opt").Id("outputs." + strcase.ToCamel(action)),
	).Block()

	err := f.Save("usecases/" + strcase.ToSnake(name) + ".go")
	if err != nil {
		return err
	}
	fmt.Printf("\x1b[33m%s\x1b[0m usecases/"+strcase.ToSnake(name)+".go\n", "create")

	f = NewFile("inputs")
	f.Type().Id(strcase.ToCamel(action)).Struct()
	err = f.Save("usecases/inputs/" + strcase.ToSnake(name) + ".go")
	if err != nil {
		return err
	}
	fmt.Printf("\x1b[33m%s\x1b[0m usecases/inputs/"+strcase.ToSnake(name)+".go\n", "create")

	f = NewFile("outputs")
	f.Type().Id(strcase.ToCamel(action)).Struct()
	err = f.Save("usecases/outputs/" + strcase.ToSnake(name) + ".go")
	if err != nil {
		return err
	}
	fmt.Printf("\x1b[33m%s\x1b[0m usecases/outputs/"+strcase.ToSnake(name)+".go\n", "create")

	return nil
}
