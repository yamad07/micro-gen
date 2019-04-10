package cmd

import (
	"flag"
	"log"

	"github.com/ispec-inc/micro-gen/lib/generator"
	"github.com/micro/cli"
)

func Scaffold(c *cli.Context) {
	flag.Parse()
	args := flag.Args()
	name := args[1]
	action := args[2]
	err := generator.Controller(name, action)
	if err != nil {
		log.Fatal(err)
	}

	err = generator.Usecase(name, action)
	if err != nil {
		log.Fatal(err)
	}
}

func Entity(c *cli.Context) {
	flag.Parse()
	args := flag.Args()
	name := args[1]
	err := generator.Entity(name)
	if err != nil {
		log.Fatal(err)
	}
}
