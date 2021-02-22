## Mokeylang

Learning by building an interpreter in golang

## Tokens


### Lexer

Take source code as input and output the tokens representing source code.
We initialize the lexer with our source code and repeatedly call next token to go through the code, token by token. Source code has type string
- we dont track file names and line numbers for now
- `NextToken()` is used to iterate through the source code

Started with creating a lexer test, so we have a sense of what we need to achieve (TDD)

### Parser

Takes in input data, and builds a data structure (AST in our case). The goal here is to give structure to the otherwise meaningless input. Here, the parser is the equivalent of `JSON.parse()` in js for json objects


### Recursive descent parsing
- `parseProgram` - entrypoint
- constructs root node of te AST (`newProgramASTNode()`)
- Builds child nodes, statements etc
- parseExpression - recursive portion

#### Misc

- The syntax consists of let statements, return statements and expressions. Pretty cool perspective to think from.
