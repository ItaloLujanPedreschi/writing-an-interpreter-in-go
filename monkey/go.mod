module main

go 1.24.1

replace monkey/lexer => ./lexer

replace monkey/parser => ./parser

replace monkey/repl => ./repl

replace monkey/token => ./token

require (
	monkey/lexer v0.0.0-00010101000000-000000000000
	monkey/parser v0.0.0-00010101000000-000000000000
	monkey/repl v0.0.0-00010101000000-000000000000
	monkey/token v0.0.0-00010101000000-000000000000
)
