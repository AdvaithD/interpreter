## Mokeylang

Learning by building an interpreter in golang

## Tokens


### Lexer

Take source code as input and output the tokens representing source code.
We initialize the lexer with our source code and repeatedly call next token to go through the code, token by token. Source code has type string
- we dont track file names and line numbers for now
- `NextToken()` is used to iterate through the source code

Started with creating a lexer test, so we have a sense of what we need to achieve (TDD)