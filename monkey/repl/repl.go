package repl

// RELP Read, Eval, Print, Lopp

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kawa1214/go-interprinter/monkey/lexer"
	"github.com/kawa1214/go-interprinter/monkey/token"
)

// PROMPT is
const PROMPT = ">> "

// Start ソースコードをトークナイズ，トークレンツを表示する
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
