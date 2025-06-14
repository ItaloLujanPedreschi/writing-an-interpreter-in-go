module monkey/repl

go 1.24.1

replace monkey/evaluator => ../evaluator

replace monkey/lexer => ../lexer

replace monkey/object => ../object

replace monkey/parser => ../parser

require (
	monkey/evaluator v0.0.0-00010101000000-000000000000
	monkey/lexer v0.0.0-00010101000000-000000000000
	monkey/object v0.0.0-00010101000000-000000000000
	monkey/parser v0.0.0-00010101000000-000000000000
)
