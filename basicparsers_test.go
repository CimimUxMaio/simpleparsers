package simpleparsers

import "testing"

var (
	digitParser       Parser = NewDigitParser()
	letterParser      Parser = NewLetterParser()
	alfaNumCharParser Parser = NewAlphanumericParser()
	wordParser        Parser = NewWordParser()
	numberParser      Parser = NewIntegerParser()
)

func TestDigitParser(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"1asdasd", NewParserOutput("1", "asdasd")},
		stringParserOutputTestCase{"5012", NewParserOutput("5", "012")},
		stringParserOutputTestCase{"hello", nil},
		stringParserOutputTestCase{"+2-%2", nil},
		stringParserOutputTestCase{"$3-%2", nil},
		stringParserOutputTestCase{"asd32", nil},
	}

	parserAssertAllEqualsParserOutput(t, digitParser, testCases)
}

func TestLetterParser(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"a1231", NewParserOutput("a", "1231")},
		stringParserOutputTestCase{"hello", NewParserOutput("h", "ello")},
		stringParserOutputTestCase{"OlLeH", NewParserOutput("O", "lLeH")},
		stringParserOutputTestCase{"12345", nil},
		stringParserOutputTestCase{"+23Abc", nil},
		stringParserOutputTestCase{"$a245", nil},
	}

	parserAssertAllEqualsParserOutput(t, letterParser, testCases)
}

func TestAlfanumericParser(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"a1231", NewParserOutput("a", "1231")},
		stringParserOutputTestCase{"hello", NewParserOutput("h", "ello")},
		stringParserOutputTestCase{"12345", NewParserOutput("1", "2345")},
		stringParserOutputTestCase{"OlLeH", NewParserOutput("O", "lLeH")},
		stringParserOutputTestCase{"+23Abc", nil},
		stringParserOutputTestCase{"$a245", nil},
	}

	parserAssertAllEqualsParserOutput(t, alfaNumCharParser, testCases)
}

func TestWordParser(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"hello world", NewParserOutput("hello", " world")},
		stringParserOutputTestCase{"a1231", NewParserOutput("a", "1231")},
		stringParserOutputTestCase{"asd123", NewParserOutput("asd", "123")},
		stringParserOutputTestCase{"IusaqE123", NewParserOutput("IusaqE", "123")},
		stringParserOutputTestCase{"dSa+DsA.", NewParserOutput("dSa", "+DsA.")},
		stringParserOutputTestCase{"12345", nil},
		stringParserOutputTestCase{"1aloha", nil},
		stringParserOutputTestCase{"+23Abc", nil},
	}

	parserAssertAllEqualsParserOutput(t, wordParser, testCases)
}

func TestNumberParser(t *testing.T) {
	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"123123", NewParserOutput("123123", "")},
		stringParserOutputTestCase{"9876asd", NewParserOutput("9876", "asd")},
		stringParserOutputTestCase{"1hello", NewParserOutput("1", "hello")},
		stringParserOutputTestCase{"+754//", nil},
	}

	parserAssertAllEqualsParserOutput(t, numberParser, testCases)
}

func TestRegexParserEquivalentToWordParser(t *testing.T) {
	regexParser, _ := NewRegexParser("[a-zA-Z]+") // Equivalent to wordParser

	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"hello world", NewParserOutput("hello", " world")},
		stringParserOutputTestCase{"a1231", NewParserOutput("a", "1231")},
		stringParserOutputTestCase{"asd123", NewParserOutput("asd", "123")},
		stringParserOutputTestCase{"IusaqE123", NewParserOutput("IusaqE", "123")},
		stringParserOutputTestCase{"dSa+DsA.", NewParserOutput("dSa", "+DsA.")},
		stringParserOutputTestCase{"12345", nil},
		stringParserOutputTestCase{"1aloha", nil},
		stringParserOutputTestCase{"+23Abc", nil},
	}

	parserAssertAllEqualsParserOutput(t, regexParser, testCases)
}

func TestRegexParserCustom(t *testing.T) {
	regexParser, _ := NewRegexParser("[3-9]+")

	var testCases []stringParserOutputTestCase = []stringParserOutputTestCase{
		stringParserOutputTestCase{"3412", NewParserOutput("34", "12")},
		stringParserOutputTestCase{"9876asd", NewParserOutput("9876", "asd")},
		stringParserOutputTestCase{"123123", nil},
		stringParserOutputTestCase{"+754//", nil},
		stringParserOutputTestCase{"1hello", nil},
		stringParserOutputTestCase{"hello world", nil},
	}

	parserAssertAllEqualsParserOutput(t, regexParser, testCases)
}
