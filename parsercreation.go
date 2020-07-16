package simpleparsers

import (
	"regexp"
)

// NewRegexParser :
// Creates a Parser that parses a string matching a prefix that matches the given regex.
// May return an error if the given reex has not a valid format.
func NewRegexParser(regexPattern string) (Parser, error) {
	regex, err := regexp.CompilePOSIX(regexPattern)
	return &regexParser{regex: regex}, err
}

// NewAnyCharParser :
// Creates a Parser that parses a string matching any given character.
func NewAnyCharParser() Parser {
	return &genericParser{parseMethod: parseAny}
}

// NewDigitParser :
// Creates a Parser that parses a string matching the first character if it is a digit according to unicode.IsDigit.
//
// It is equivalent to:
//	Conditional(NewAnyCharParser(), startsWithDigit)
// where `startsWithDigit` returns true if the head of the string is in fact a _digit_.
func NewDigitParser() Parser {
	return Conditional(NewAnyCharParser(), startsWithDigit)
}

// NewLetterParser :
// Creates a Parser that parses a string matching the first character if it is a letter according to unicode.IsLetter.
//
// It is equivalent to:
// 	Conditional(NewAnyCharParser(), startsWithLetter)
// where `startsWithLetter` returns true if the head of the string is in fact a _letter_.
func NewLetterParser() Parser {
	return Conditional(NewAnyCharParser(), startsWithLetter)
}

// NewAlphanumericParser :
// Creates a Parser that parses a string matching the first character if it is either a _letter_ or a _digit_ according to unicode.IsLetter and unicode.IsDigit.
//
// It is equivalent to:
//	Either(NewLetterParser(), NewDigitParser())
func NewAlphanumericParser() Parser {
	return Either(NewLetterParser(), NewDigitParser())
}

// NewIntegerParser :
// Creates a Parser that parses a string matching the first integer number (sequence of _digits_).
//
// It is equivalent to:
//	KleenePlus(NewDigitParser())
func NewIntegerParser() Parser {
	return KleenePlus(NewDigitParser())
}

// NewWordParser :
// Creates a Parser that parses a string matching the first word (sequence of _letters_).
//
// It is equivalent to:
// 	KleenePlus(NewLetterParser())
func NewWordParser() Parser {
	return KleenePlus(NewLetterParser())
}

// NewCharParser :
// Creates a Parser that parses a string matching the first character only if it is the same as the given character.
func NewCharParser(char rune) Parser {
	return Conditional(NewAnyCharParser(), func(match string) bool { return startsWithChar(match, char) })
}

// NewNumberParser :
// Creates a Parser that parses a string matching the first number (either an integer or a floating point number).
//
// It is equivalent to:
//	Sequence(NewIntegerParser(), Optional(Sequence(NewCharParser('.'), NewIntegerParser())))
func NewNumberParser() Parser {
	return Sequence(NewIntegerParser(), Optional(Sequence(NewCharParser('.'), NewIntegerParser())))
}
