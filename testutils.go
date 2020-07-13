package simpleparsers

import "testing"

type stringParserOutputTestCase struct {
	input          string
	expectedOutput *ParserOutput
}

func assertEqualsParserOutput(t *testing.T, input string, expectedOutput *ParserOutput, actualOutput *ParserOutput) {
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

func parserAssertAllEqualsParserOutput(t *testing.T, parser Parser, testCases []stringParserOutputTestCase) {
	for _, testCase := range testCases {
		poutput := parser.Parse(testCase.input)
		assertEqualsParserOutput(t, testCase.input, testCase.expectedOutput, poutput)
	}
}
