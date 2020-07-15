package simpleparsers

// Sequence :
// Given two parsers, returns a new one that parses as the two parsers consecutively, using the remainder of the first one as input for the second.
// If for a certain input either parser fails, the resulting parser of sequencing both will also fail for this input.
func Sequence(parser1 Parser, parser2 Parser) Parser {
	return &genericParser{
		parseMethod: func(input string) (*ParserOutput, error) {
			return parseConsecutively(parser1, parser2, input)
		},
	}
}

// Either :
// Given two parsers, returns a new one that parses a certain input as the first one. If it fails, parses as the second.
// If both fail then this parser will also fail.
func Either(parser1 Parser, parser2 Parser) Parser {
	return &genericParser{
		parseMethod: func(input string) (*ParserOutput, error) {
			return parseWithEither(parser1, parser2, input)
		},
	}
}

// KleenePlus :
// Returns a parser that parses a certain input as the given parser iteratively until there is no more matches.
// If there is no matches, the parser returns an error.
func KleenePlus(parser Parser) Parser {
	return &genericParser{
		parseMethod: func(input string) (*ParserOutput, error) {
			return parseIterativelyAtLeastOnce(parser, input)
		},
	}
}

// Optional :
// Returns a parser that parses a certain input as the given parser iteratively until there is no more matches.
// If there is no matches, the parser returns an error.
func Optional(parser Parser) Parser {
	return &genericParser{
		parseMethod: func(input string) (*ParserOutput, error) {
			return parseOptionaly(input, parser)
		},
	}
}

// KleeneStar :
// Returns a parser that parses a certain input as the given parser iteratively until there is no more matches.
// If there is no matches, the parser matches the empty string ("").
func KleeneStar(parser Parser) Parser {
	return Optional(KleenePlus(parser))
}
