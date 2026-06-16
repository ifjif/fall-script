package cmd

import (
	"fmt"
	"os"
	"strings"

	"zzc/fall-script/src/compiler"
	"zzc/fall-script/src/evaluator"
	"zzc/fall-script/src/macro"
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
	data, err := os.ReadFile(c.File)
	if err != nil {
		panic(err)
	}

	input := string(data)

	p := parser.NewParser(input)
	program := p.Parse()
	if len(p.Errors()) != 0 {
		error := strings.Join(p.Errors(), "\n")
		fmt.Println(error)
	}

	comp := compiler.NewCompiler(program)
	comp.Compile()

	if c.Dump {
		data := comp.Dump()
		fmt.Println("DUMP-------------------------------------------")
		fmt.Printf("%v\n", data)
		outName := c.File + "o"
		os.WriteFile(outName, data, 0o644)
		fmt.Println("UNDUMP-------------------------------------------")
		cf := comp.Undump(data)
		fmt.Printf("%+v\n", cf)

		fsVM := vm.NewFsVM(cf)
		fsVM.Run()
	} else {
		fsVM := vm.NewFsVM(comp.MainFn())

		fsVM.Run()
	}
}
