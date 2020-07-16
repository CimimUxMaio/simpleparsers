package simpleparsers

import (
	"errors"
	"reflect"
	"runtime"
)

func parseAny(input string) (*ParserOutput, error) {
	if len(input) == 0 {
		return nil, errors.New("no character match for empty string")
	}

	match := string(input[0])
	remainder := ""

	if len(input) > 1 {
		remainder = input[1:]
	}

	return &ParserOutput{Match: match, Remainder: remainder}, nil
}

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

func parseWithCondition(input string, parser Parser, condition func(string) bool) (*ParserOutput, error) {
	output, err := parser.Parse(input)

	if err != nil {
		return nil, err
	}

	if !condition(output.Match) {
		errNoMatch := errors.New("no match found that satisfies the condition: " + runtime.FuncForPC(reflect.ValueOf(condition).Pointer()).Name() + " with input: \"" + input + "\"")
		return nil, errNoMatch
	}

	return output, nil
}

func parseOptionaly(input string, parser Parser) (*ParserOutput, error) {
	output, err := parser.Parse(input)

	if err != nil {
		output = &ParserOutput{Match: "", Remainder: input}
	}

	return output, nil
}

func parseAndIgnoreMatch(input string, parser Parser) (*ParserOutput, error) {
	output, err := parser.Parse(input)

	if err != nil {
		return nil, err
	}

	return &ParserOutput{Match: "", Remainder: output.Remainder}, nil
}
