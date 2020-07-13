package simpleparsers

import (
	"strings"
	"unicode"
)

func parseConsecutively(parser1 Parser, parser2 Parser, input string) *ParserOutput {
	output1 := parser1.Parse(input)

	if output1 == nil {
		return nil
	}

	output2 := parser2.Parse(output1.Remainder)

	if output2 == nil {
		return nil
	}

	return &ParserOutput{Match: output1.Match + output2.Match, Remainder: output2.Remainder}
}

func parseWithEither(parser1 Parser, parser2 Parser, input string) (output *ParserOutput) {
	output = parser1.Parse(input)

	if output == nil {
		output = parser2.Parse(input)
	}

	return output
}

func parseLonguest(parser Parser, input string) *ParserOutput {
	longuestMatch := ""
	remainder := ""

	output := parser.Parse(input)

	if output == nil {
		return nil
	}

	for output != nil {
		longuestMatch += output.Match
		remainder = output.Remainder
		output = parser.Parse(remainder)
	}

	return &ParserOutput{Match: longuestMatch, Remainder: remainder}
}

func parseWithCondition(input string, condition func(r rune) bool) (output *ParserOutput) {

	output = nil
	if len(input) > 0 {
		if head := rune(input[0]); condition(head) {
			var match string = string(head)
			output = NewParserOutput(match, strings.TrimPrefix(input, match))
		}
	}

	return output
}

func parseDigit(input string) *ParserOutput {
	return parseWithCondition(input, unicode.IsDigit)
}

func parseLetter(input string) *ParserOutput {
	return parseWithCondition(input, unicode.IsLetter)
}
