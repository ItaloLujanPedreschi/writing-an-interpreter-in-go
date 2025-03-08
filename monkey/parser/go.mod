module monkey/parser

go 1.24.1

replace monkey/ast => ../ast

replace monkey/lexer => ../lexer

replace monkey/token => ../token

require (
	monkey/ast v0.0.0-00010101000000-000000000000
	monkey/lexer v0.0.0-00010101000000-000000000000
	monkey/token v0.0.0-00010101000000-000000000000
)
