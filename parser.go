package simpleparsers

import (
	"regexp"
	"strings"
)

// Parser - can parse a string, returning a *ParserOutput that wraps the matching value and the remainder,
// so that the matching value + the remainder equals the input. If there is no match, it returns nil.
type Parser interface {
	Parse(input string) *ParserOutput
}

// ParserOutput - the output type of Parser. Wraps the matching value and the remainder.
type ParserOutput struct {
	Match     string
	Remainder string
}

// NewParserOutput - returns a pointer to a ParserOutput value with the given data.
func NewParserOutput(match, remainder string) *ParserOutput {
	return &ParserOutput{match, remainder}
}

// AsString - formats a *ParserOutput to a string.
func (poutput *ParserOutput) AsString() string {
	return "{ Match: " + poutput.Match + ", Remainder: " + poutput.Remainder + " }"
}

// Equals - Compares two *ParserOutput values by their match and remainder.
func (poutput *ParserOutput) Equals(otherpoutput *ParserOutput) bool {
	return poutput.Match == otherpoutput.Match && poutput.Remainder == otherpoutput.Remainder
}

type genericParser struct {
	parseMethod func(string) *ParserOutput
}

func (gp *genericParser) Parse(input string) *ParserOutput {
	return gp.parseMethod(input)
}

type regexParser struct {
	regex *regexp.Regexp
}

func (rp *regexParser) Parse(input string) *ParserOutput {
	matchLocation := rp.regex.FindStringIndex(input)

	if matchLocation == nil || matchLocation[0] != 0 {
		return nil
	}

	match := input[matchLocation[0]:matchLocation[1]]
	remainder := strings.TrimPrefix(input, match)
	return &ParserOutput{Match: match, Remainder: remainder}
}
