package simpleparsers

import (
	"errors"
	"reflect"
	"runtime"
	"strings"
	"unicode"
)

func parseConsecutively(parser1 Parser, parser2 Parser, input string) (*ParserOutput, error) {
	output1, err1 := parser1.Parse(input)

	if err1 != nil {
		return nil, err1
	}

	output2, err2 := parser2.Parse(output1.Remainder)

	if err2 != nil {
		return nil, err2
	}

	return &ParserOutput{Match: output1.Match + output2.Match, Remainder: output2.Remainder}, nil
}

func parseWithEither(parser1 Parser, parser2 Parser, input string) (*ParserOutput, error) {
	output, err := parser1.Parse(input)

	if err != nil {
		output, err = parser2.Parse(input)
	}

	return output, err
}

func parseIterativelyAtLeastOnce(parser Parser, input string) (*ParserOutput, error) {
	longuestMatch := ""
	remainder := ""

	output, err := parser.Parse(input)

	if err != nil {
		return nil, err
	}

	for err == nil {
		longuestMatch += output.Match
		remainder = output.Remainder
		output, err = parser.Parse(remainder)
	}

	return &ParserOutput{Match: longuestMatch, Remainder: remainder}, nil
}

func parseWithCondition(input string, condition func(r rune) bool) (*ParserOutput, error) {
	errNoMatch := errors.New("No match found for character with condition: " + runtime.FuncForPC(reflect.ValueOf(condition).Pointer()).Name() + " and input: \"" + input + "\"")
	if len(input) == 0 {
		return nil, errNoMatch
	}

	head := rune(input[0])

	if !condition(head) {
		return nil, errNoMatch
	}

	match := string(head)
	output := &ParserOutput{Match: match, Remainder: strings.TrimPrefix(input, match)}

	return output, nil
}

func parseDigit(input string) (*ParserOutput, error) {
	return parseWithCondition(input, unicode.IsDigit)
}

func parseLetter(input string) (*ParserOutput, error) {
	return parseWithCondition(input, unicode.IsLetter)
}
