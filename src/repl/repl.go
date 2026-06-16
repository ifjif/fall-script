package repl

import (
	"bufio"
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

const PROMPT = ">>"

func Repl() {
	scanner := bufio.NewReader(os.Stdin)
	env := object.NewEnvironment()
	for {
		fmt.Print(PROMPT)
		input, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		p := parser.NewParser(input)
		program := p.Parse()
		if len(p.Errors()) != 0 {
			error := strings.Join(p.Errors(), "\n")
			fmt.Println(error)
			continue
		}
		// fmt.Println(program.String())

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
}

func ReplVM() {
	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(PROMPT)
		input, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		p := parser.NewParser(input)
		program := p.Parse()
		if len(p.Errors()) != 0 {
			error := strings.Join(p.Errors(), "\n")
			fmt.Println(error)
			continue
		}
		// fmt.Println(program.String())

		c := compiler.NewCompiler(program)
		c.Compile()
		//	fmt.Println(c.CurrentInstructions())
		//	consts := c.CurrentConstant()
		//	printFn(consts)
		fsVM := vm.NewFsVM(c.MainFn())
		fsVM.Run()
	}
}

func printFn(consts []object.Object) {
	for _, ct := range consts {
		switch ct := ct.(type) {
		case *object.CompiledFunction:
			fmt.Println(ct.Instructions.String())
			printFn(ct.Constants)
		}
	}
}
