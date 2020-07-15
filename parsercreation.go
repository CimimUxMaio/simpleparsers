package simpleparsers

import (
	"regexp"
)

// NewRegexParser - creates a Parser that parses a string matching a prefix that matches the given regex.
// May return an error if the given reex has not a valid format.
func NewRegexParser(regexPattern string) (Parser, error) {
	regex, err := regexp.CompilePOSIX(regexPattern)
	return &regexParser{regex: regex}, err
}

// NewDigitParser - creates a Parser that parses a string matching the first character if it is a digit according to unicode.IsDigit.
func NewDigitParser() Parser {
	return &genericParser{parseMethod: parseDigit}
}

// NewLetterParser - creates a Parser that parses a string matching the first character if it is a letter according to unicode.IsLetter.
func NewLetterParser() Parser {
	return &genericParser{parseMethod: parseLetter}
}

// NewAlphanumericParser - creates a Parser that parses a string matching the first character if it is either a letter or a digit according to unicode.IsLetter and unicode.IsDigit.
//
// It is equivalent to:
//	Either(NewLetterParser(), NewDigitParser())
func NewAlphanumericParser() Parser {
	return Either(NewLetterParser(), NewDigitParser())
}

// NewIntegerParser - creates a Parser that parses a string matching the first integer number (sequence of digits).
//
// It is equivalent to:
//	KleenePlus(NewDigitParser())
func NewIntegerParser() Parser {
	return KleenePlus(NewDigitParser())
}

// NewWordParser - creates a Parser that parses a string matching the first word (sequence of letters).
//
// It is equivalent to:
//	KleenePlus(NewLetterParser())
func NewWordParser() Parser {
	return KleenePlus(NewLetterParser())
}

// NewCharParser - creates a Parser that parses a string matching the first character only if it is the same as the given character.
func NewCharParser(char rune) Parser {
	return &genericParser{
		parseMethod: func(input string) (*ParserOutput, error) {
			return parseWithCondition(input, func(match rune) bool { return match == char })
		},
	}
}

// NewNumberParser - creates a Parser that parses a string matching the first number (either an integer or a floating point number).
//
// It is equivalent to:
//	Sequence(NewIntegerParser(), Optional(Sequence(NewCharParser('.'), NewIntegerParser())))
func NewNumberParser() Parser {
	return Sequence(NewIntegerParser(), Optional(Sequence(NewCharParser('.'), NewIntegerParser())))
}
