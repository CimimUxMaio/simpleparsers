package tests

import (
	"testing"

	"github.com/cimimuxmaio/simpleparsers"
)

type stringParserOutputTestCase struct {
	input          string
	expectedOutput *simpleparsers.ParserOutput
}

func newParserOutput(match, remainder string) *simpleparsers.ParserOutput {
	return &simpleparsers.ParserOutput{Match: match, Remainder: remainder}
}

func assertEqualsParserOutput(t *testing.T, input string, expectedOutput *simpleparsers.ParserOutput, actualOutput *simpleparsers.ParserOutput) {
	if expectedOutput == nil || actualOutput == nil {
		if expectedOutput != actualOutput {
			if expectedOutput == nil {
				throwTestFail(t, input, nil, actualOutput.AsString())
			} else {
				throwTestFail(t, input, expectedOutput.AsString(), nil)
			}
		}

		return
	}

	if !expectedOutput.Equals(actualOutput) {
		throwTestFail(t, input, expectedOutput.AsString(), actualOutput.AsString())
	}
}

func throwTestFail(t *testing.T, input interface{}, expected interface{}, actual interface{}) {
	t.Errorf("For input: \"%s\", the expected value was: \"%s\" but got: \"%s\" instead.", input, expected, actual)
}

func parserAssertAllEqualsParserOutput(t *testing.T, parser simpleparsers.Parser, testCases []stringParserOutputTestCase) {
	for _, testCase := range testCases {
		poutput, _ := parser.Parse(testCase.input)
		assertEqualsParserOutput(t, testCase.input, testCase.expectedOutput, poutput)
	}
}
