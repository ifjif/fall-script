package cmd

import (
	"fmt"
	"os"
	"strings"

	"zzc/fall-script/src/evaluator"
	"zzc/fall-script/src/macro"
	"zzc/fall-script/src/module"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/parser"
	"zzc/fall-script/src/vm"
)

type Cmd struct {
	File   string
	Engine string
	Dump   bool
}

func (c *Cmd) Execute() {
	data, err := os.ReadFile(c.File)
	if err != nil {
		panic(err)
	}

	input := string(data)

	env := object.NewEnvironment()
	p := parser.NewParser(input)
	program := p.Parse()
	if len(p.Errors()) != 0 {
		error := strings.Join(p.Errors(), "\n")
		fmt.Println(error)
	}

	macroEnv := object.NewEnvironment()
	macro.DefineMacros(program, macroEnv)
	np := macro.ExpandMacros(program, macroEnv)

	eval := evaluator.NewEvaluator(np, env)
	result := eval.Evaluate()

	if err, ok := result.(*object.ErrorObj); ok {
		fmt.Println(err.Inspect())
	}
}

func (c *Cmd) Interprete() {
	if c.Dump {
		loader := module.NewLoader()
		file := module.ResolveImportPath(".", c.File)
		loader.DumpFile(file)
	} else {
		fsVM := vm.NewFsVMWithFile(c.File)
		fsVM.Run()
	}
}
