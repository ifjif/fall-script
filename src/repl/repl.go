package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/evaluator"
	"zzc/fall-script/src/macro"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/parser"
	"zzc/fall-script/src/vm"
)

const PROMPT = ">>"

func Repl() {
	scanner := bufio.NewReader(os.Stdin)
	// env := object.NewEnvironment()
	for {
		fmt.Print(PROMPT)
		input, err := scanner.ReadBytes('\n')
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		p := parser.NewParser(string(input))
		program := p.Parse()
		if len(p.Errors()) != 0 {
			error := strings.Join(p.Errors(), "\n")
			fmt.Println(error)
			continue
		}
		fmt.Println(program.String())

		// doEvaluate(program, env)
	}
}

func doEvaluate(program *ast.Program, env *object.Environment) {
	macroEnv := object.NewEnvironment()
	macro.DefineMacros(program, macroEnv)
	fmt.Println("删除宏后的AST:-------------")
	fmt.Println(program.String())
	fmt.Println("-------------")
	fmt.Println("宏函数:-------------")
	fmt.Println(macroEnv.Inspect())

	nProgram := macro.ExpandMacros(program, macroEnv)
	fmt.Println("-------------")
	fmt.Println("宏展开后的AST:-------------")
	fmt.Println(nProgram.String())
	fmt.Println("-------------")

	eval := evaluator.NewEvaluator(nProgram, env)
	obj := eval.Evaluate()
	if obj != nil {
		fmt.Println(obj.Inspect())
	}
}

func ReplVM() {
	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(PROMPT)
		input, err := scanner.ReadBytes('\n')
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fsVM := vm.NewFsVMWithText(input)
		fsVM.Run()
	}
}
