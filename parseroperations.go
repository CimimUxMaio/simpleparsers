package simpleparsers

// Sequence - given two parsers, returns a new one that is equivalent to parsing with
// parser1 and parser2 consecutively.
func Sequence(parser1 Parser, parser2 Parser) Parser {
	return &genericParser{
		parseMethod: func(input string) *ParserOutput {
			return parseConsecutively(parser1, parser2, input)
		},
	}
}

// Either - given two parsers, returns a parser that returns the result of parsing with parser1.
// If it does not match anything, returns the result of parsing with parser2.
func Either(parser1 Parser, parser2 Parser) Parser {
	return &genericParser{
		parseMethod: func(input string) *ParserOutput {
			return parseWithEither(parser1, parser2, input)
		},
	}
}

// KleenePlus - returns a parser that parses with the given parser consecutively until there is no more matches.
// If there is no matches returns nil.
func KleenePlus(parser Parser) Parser {
	return &genericParser{
		parseMethod: func(input string) *ParserOutput {
			return parseLonguest(parser, input)
		},
	}
}