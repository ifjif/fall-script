package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"zzc/fall-script/src/evaluator"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/parser"
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

		eval := evaluator.NewEvaluator(program, env)
		obj := eval.Evaluate()
		if obj != nil {
			fmt.Println(obj.Inspect())
		}
	}
}
