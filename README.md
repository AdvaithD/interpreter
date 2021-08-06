## **MonkeyLang - An interpreter built in go**

Attempting to learn go by building an interpreter in golang.

### **Instructions**

1. Run `go build main.go` to build the binary or run `go run main.go`.
2. This should open up a REPL which allows you to enter MonkeyLang statements.

### **AST**
- Contains code that helps build an abstract syntax tree used by the parser.
- The root node of the AST is defined by the type `Program` which contains a slice of statements (i.e: `[]Statements`).
- Some important types are: `Statement`, `LetStatement`, `ReturnStatement`, `ExpressionStatement`, `PrefixExpression`.
- Most statements have the following functions defined as part of their interface: `TokenLiteral()`, `String()`.

### **Tokens**
- The `token` package contains a dictionary of supported keywords (e.g: `fn`, `let`, `true`, `false` etc).
- `token.LookupIdent()` retrieves the keyword given a string, else returns the type `IDENT`.


### **Lexer**

Take source code as input and output the tokens representing source code.
We initialize the lexer with our source code and repeatedly call next token to go through the code, token by token. Source code has type string.
- We dont track file names and line numbers for now. (possible future feature)
- `NextToken()` is used to iterate through the source code.

Started with creating a lexer test, so we have a sense of what we need to achieve (TDD)

### **Parser**

Takes in input data, and builds a data structure (AST in our case). The goal here is to give structure to the otherwise meaningless input. Here, the parser is the equivalent of `JSON.parse()` in js for json objects

### **Recursive descent parsing**
- `parseProgram` - entrypoint
- constructs root node of te AST (`newProgramASTNode()`)
- Builds child nodes, statements etc
- parseExpression - recursive portion

### **Pratt Parsing - Top Down Operator Precedence**

[Top Down Operator Precedence - Vaughan R. Pratt, 1973](http://tdop.github.io/)

- Each token type can have two parsing functions associated with it (infix or prefix)
- The core of pratt parsing in the context of this repo lies in the `parseExpression()` function. It ensures that the ASTs are nested correctly.
- 

An illustration of how we want our `AST` and `Parser` to work and create an optimal number of AST nodes.

Code: `1 + 2 + 3`
We want this to be represented as ((1+2) + 3) where `1+2` represents an infix expression, and `3` is an integer

```
                            *ast.InfixExpression
                            /                     \
                *ast.InfixExpression             *ast.IntegerLiteral
                /                     \                              \
    *ast.IntegerLiteral - 1        *ast.IntegerLiteral - 2            3
```

prefix e.g: `--5`
infix e.g: `2 * 2`
postfix e.g: `variableName++`
#### **Misc**

- The syntax consists of let statements, return statements and expressions. Pretty cool perspective to think from.
