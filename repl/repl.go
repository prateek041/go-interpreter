package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/prateek041/go-interpreter/lexer"
)

const PROMPT = ">"

func Start(out io.Writer, in io.Reader) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, PROMPT)
		isScanned := scanner.Scan()
		if !isScanned {
			return
		}
		input := scanner.Text()
		l := lexer.New(input)
		l.SpitTokens(out)
	}
}
