[![Build Status](https://travis-ci.org/CimimUxMaio/simpleparsers.svg?branch=master)](https://travis-ci.org/CimimUxMaio/simpleparsers)
[![BCH compliance](https://bettercodehub.com/edge/badge/CimimUxMaio/simpleparsers?branch=master)](https://bettercodehub.com/)

# simpleparsers

A package with functions to create and combine parsers.

---

## Built-in parsers

- [RegexParser](###RegexParser)
- [DigitParser](###DigitParser)
- [LetterParser](###LetterParser)
- [AlphanumericParser](###AlphanumericParser)
- [CharParser](###CharParser)
- [WordParser](###WordParser)
- [IntegerParser](###IntegerParser)
- [NumberParser](###NumberParser)

## Parser operations

- [Sequence/2](###Sequence)
- [Either/2](###Either)
- [KleenePlus/1](###KleenePlus)
- [KleeneStar/1](###KleeneStar)
- [Optional/1](###Optional)

---

## Built-in parsers

### RegexParser

Created with `NewRegexParser(<regex>)`. Parses a string matching a prefix that matches the given regex.

`NewRegexParser/2` may return an error if the given regex has not a valid format.

##### Examples:

**TODO**

### DigitParser

Created with `NewDigitParser()`. Parses a string matching the first character if it is a _digit_ according to `unicode.IsDigit`.

##### Examples:

**TODO**

### LetterParser

Created with `NewLetterParser()`. Parses a string matching the first character if it is a _letter_ according to `unicode.IsLetter`.

##### Examples:

**TODO**

### AlphanumericParser

Created with `NewAlphanumericParser()`. Parses a string matching the first character if it is either a _letter_ or a _digit_ according to `unicode.IsLetter` and `unicode.IsDigit`.

It is equivalent to: [`Either(NewLetterParser(), NewDigitParser())`](###Either)

##### Examples:

**TODO**

### CharParser

Created with `NewCharParser(<a_character>)`. Parses a string matching the first character only if it is the same as the given character.

##### Examples:

**TODO**

### WordParser

Created with `NewLetterParser()`. Parses a string matching the first word (sequence of _letters_).

It is equivalent to: [`KleenePlus(NewLetterParser())`](###KleenePlus)

##### Examples:

**TODO**

### IntegerParser

Created with `NewIntegerParser()`. Parses a string matching the first integer number (sequence of _digits_).

It is equivalent to: [`KleenePlus(NewDigitParser())`](###KleenePlus)

##### Examples:

**TODO**

### NumberParser

Created with `NewNumberParser()`. Parses a string matching the first **number** (either an _integer_ or a _floating point number_).

It is equivalent to:

`Sequence(NewIntegerParser(), Optional(Sequence(NewCharParser('.'), NewIntegerParser())))`

##### See:

- [Sequence/1](###Sequence)
- [Optional/1](###Optional)

##### Examples:

**TODO**

---

## Parser operations

### Sequence

Given two parsers, returns a new one that parses as the two parsers consecutively, using the remainder of the first one as input for the second.

If for a certain input either parser fails, the resulting parser of sequencing both will also fail for this input.

##### Examples:

**TODO**

### Either

Given two parsers, returns a new one that parses a certain input as the first one. If it fails, parses as the second. If both fail then this parser will also fail.

##### Examples:

**TODO**

### KleenePlus

Returns a parser that parses a certain input as the given parser iteratively until there is no more matches.

If there is **no** matches, the parser returns an error.

##### Examples:

**TODO**

### KleeneStar

Returns a parser that parses a certain input as the given parser iteratively until there is no more matches.

If there is **no** matches, the parser matches the empty string (`""`).

##### Examples:

**TODO**

### Optional

Returns a parser that parses a certain input as the given parser iteratively until there is no more matches.

If there is **no** matches, the parser returns an error.

##### Examples:

**TODO**
