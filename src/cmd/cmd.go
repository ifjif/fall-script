package cmd

import (
	"fmt"
	"os"
	"strings"

	"zzc/fall-script/src/evaluator"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/parser"
)

type Cmd struct {
	File string
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

	eval := evaluator.NewEvaluator(program, env)
	obj := eval.Evaluate()

	if obj != nil {
		fmt.Println(obj.Inspect())
	}
}
