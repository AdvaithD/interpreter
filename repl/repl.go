package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkeylang/lexer"
	"monkeylang/token"
)

const PROMPT = "☄ | "

// Start - read from input source until you find a newline
// take the line, pass it to a lexer instance, print all tokens lexer returns until EOF
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
