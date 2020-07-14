package simpleparsers

import (
	"regexp"
)

// NewRegexParser - with the given regex pattern creates a Parser that matches
// strings that match the regex. May return an error if the regex pattern received
// doesn't correspond to a valid regex format.
func NewRegexParser(regexPattern string) (Parser, error) {
	regex, err := regexp.CompilePOSIX(regexPattern)
	return &regexParser{regex: regex}, err
}

// NewDigitParser - creates a Parser that matches single characters that represent a digit (unicode.IsDigit).
func NewDigitParser() Parser {
	return &genericParser{parseMethod: parseDigit}
}

// NewLetterParser -creates a Parser that matches single characters that represent a letter (unicode.IsLetter).
func NewLetterParser() Parser {
	return &genericParser{parseMethod: parseLetter}
}

// NewAlphanumericParser - creates a Parser that matches with either a digit or a letter.
func NewAlphanumericParser() Parser {
	return Either(NewLetterParser(), NewDigitParser())
}

// NewIntegerParser - creates a Parser that matches with a sequence of digits.
func NewIntegerParser() Parser {
	return KleenePlus(NewDigitParser())
}

// NewWordParser - creates a Parser that matches with a sequence of letters (word).
func NewWordParser() Parser {
	return KleenePlus(NewLetterParser())
}

// NewCharParser - creates a Parser that matches only with the given character.
func NewCharParser(char rune) Parser {
	return &genericParser{
		parseMethod: func(input string) (*ParserOutput, error) {
			return parseWithCondition(input, func(match rune) bool { return match == char })
		},
	}
}

// NewNumberParser - creates a Parser that matches with every kind of number (integer or floating point).
func NewNumberParser() Parser {
	return Sequence(NewIntegerParser(), Optional(Sequence(NewCharParser('.'), NewIntegerParser())))
}
