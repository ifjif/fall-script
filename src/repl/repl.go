package repl

import (
	"bufio"
	"fmt"
	"os"

	"zzc/fall-script/src/lexer"
	"zzc/fall-script/src/token"
)

const PROMPT = ">>"

func Repl() {
	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(PROMPT)
		input, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Printf("%v\n", err)
		}

		l := lexer.NewLexer(input)
		tok := l.NextToken()
		for tok.Type != token.EOF {
			fmt.Printf("%+v\n", tok)
			tok = l.NextToken()
		}
	}
}
