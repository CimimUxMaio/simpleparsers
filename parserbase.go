package simpleparsers

import (
	"errors"
	"regexp"
	"strings"
)

// Parser - can parse a string, returning a *ParserOutput that wraps the matching value and the remainder,
// so that the matching value + the remainder equals the input. If there is no match, it returns (nil, error).
type Parser interface {
	Parse(input string) (*ParserOutput, error)
}

// ParserOutput - the output type of Parser. Wraps the matching value and the remainder.
type ParserOutput struct {
	Match     string
	Remainder string
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
	parseMethod func(string) (*ParserOutput, error)
}

func (gp *genericParser) Parse(input string) (*ParserOutput, error) {
	return gp.parseMethod(input)
}

type regexParser struct {
	regex *regexp.Regexp
}

func (rp *regexParser) Parse(input string) (*ParserOutput, error) {
	matchLocation := rp.regex.FindStringIndex(input)

	if matchLocation == nil || matchLocation[0] != 0 {
		return nil, errors.New("no match found for regex: \"" + rp.regex.String() + "\" and input: \"" + input + "\"")
	}

	match := input[matchLocation[0]:matchLocation[1]]
	remainder := strings.TrimPrefix(input, match)
	return &ParserOutput{Match: match, Remainder: remainder}, nil
}
