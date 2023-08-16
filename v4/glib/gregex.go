// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// Specifies the type of the function passed to g_regex_replace_eval().
// It is called for each occurrence of the pattern in the string passed
// to g_regex_replace_eval(), and it should append the replacement to
// @result.
type RegexEvalCallback func(*MatchInfo, *String, uintptr) bool

// A GMatchInfo is an opaque struct used to return information about
// matches.
type MatchInfo struct {
}

// The g_regex_*() functions implement regular
// expression pattern matching using syntax and semantics similar to
// Perl regular expression.
//
// Some functions accept a @start_position argument, setting it differs
// from just passing over a shortened string and setting %G_REGEX_MATCH_NOTBOL
// in the case of a pattern that begins with any kind of lookbehind assertion.
// For example, consider the pattern "\Biss\B" which finds occurrences of "iss"
// in the middle of words. ("\B" matches only if the current position in the
// subject is not a word boundary.) When applied to the string "Mississipi"
// from the fourth byte, namely "issipi", it does not match, because "\B" is
// always false at the start of the subject, which is deemed to be a word
// boundary. However, if the entire string is passed , but with
// @start_position set to 4, it finds the second occurrence of "iss" because
// it is able to look behind the starting point to discover that it is
// preceded by a letter.
//
// Note that, unless you set the %G_REGEX_RAW flag, all the strings passed
// to these functions must be encoded in UTF-8. The lengths and the positions
// inside the strings are in bytes and not in characters, so, for instance,
// "\xc3\xa0" (i.e. "à") is two bytes long but it is treated as a
// single character. If you set %G_REGEX_RAW the strings can be non-valid
// UTF-8 strings and a byte is treated as a character, so "\xc3\xa0" is two
// bytes and two characters long.
//
// When matching a pattern, "\n" matches only against a "\n" character in
// the string, and "\r" matches only a "\r" character. To match any newline
// sequence use "\R". This particular group matches either the two-character
// sequence CR + LF ("\r\n"), or one of the single characters LF (linefeed,
// U+000A, "\n"), VT vertical tab, U+000B, "\v"), FF (formfeed, U+000C, "\f"),
// CR (carriage return, U+000D, "\r"), NEL (next line, U+0085), LS (line
// separator, U+2028), or PS (paragraph separator, U+2029).
//
// The behaviour of the dot, circumflex, and dollar metacharacters are
// affected by newline characters, the default is to recognize any newline
// character (the same characters recognized by "\R"). This can be changed
// with %G_REGEX_NEWLINE_CR, %G_REGEX_NEWLINE_LF and %G_REGEX_NEWLINE_CRLF
// compile options, and with %G_REGEX_MATCH_NEWLINE_ANY,
// %G_REGEX_MATCH_NEWLINE_CR, %G_REGEX_MATCH_NEWLINE_LF and
// %G_REGEX_MATCH_NEWLINE_CRLF match options. These settings are also
// relevant when compiling a pattern if %G_REGEX_EXTENDED is set, and an
// unescaped "#" outside a character class is encountered. This indicates
// a comment that lasts until after the next newline.
//
// When setting the %G_REGEX_JAVASCRIPT_COMPAT flag, pattern syntax and pattern
// matching is changed to be compatible with the way that regular expressions
// work in JavaScript. More precisely, a lonely ']' character in the pattern
// is a syntax error; the '\x' escape only allows 0 to 2 hexadecimal digits, and
// you must use the '\u' escape sequence with 4 hex digits to specify a unicode
// codepoint instead of '\x' or 'x{....}'. If '\x' or '\u' are not followed by
// the specified number of hex digits, they match 'x' and 'u' literally; also
// '\U' always matches 'U' instead of being an error in the pattern. Finally,
// pattern matching is modified so that back references to an unset subpattern
// group produces a match with the empty string instead of an error. See
// pcreapi(3) for more information.
//
// Creating and manipulating the same #GRegex structure from different
// threads is not a problem as #GRegex does not modify its internal
// state between creation and destruction, on the other hand #GMatchInfo
// is not threadsafe.
//
// The regular expressions low-level functionalities are obtained through
// the excellent
// [PCRE](http://www.pcre.org/)
// library written by Philip Hazel.
type Regex struct {
}

// Flags specifying compile-time options.
type RegexCompileFlags int

const (

	// Letters in the pattern match both upper- and
	//     lowercase letters. This option can be changed within a pattern
	//     by a "(?i)" option setting.
	GRegexCaselessValue RegexCompileFlags = 1
	// By default, GRegex treats the strings as consisting
	//     of a single line of characters (even if it actually contains
	//     newlines). The "start of line" metacharacter ("^") matches only
	//     at the start of the string, while the "end of line" metacharacter
	//     ("$") matches only at the end of the string, or before a terminating
	//     newline (unless %G_REGEX_DOLLAR_ENDONLY is set). When
	//     %G_REGEX_MULTILINE is set, the "start of line" and "end of line"
	//     constructs match immediately following or immediately before any
	//     newline in the string, respectively, as well as at the very start
	//     and end. This can be changed within a pattern by a "(?m)" option
	//     setting.
	GRegexMultilineValue RegexCompileFlags = 2
	// A dot metacharacter (".") in the pattern matches all
	//     characters, including newlines. Without it, newlines are excluded.
	//     This option can be changed within a pattern by a ("?s") option setting.
	GRegexDotallValue RegexCompileFlags = 4
	// Whitespace data characters in the pattern are
	//     totally ignored except when escaped or inside a character class.
	//     Whitespace does not include the VT character (code 11). In addition,
	//     characters between an unescaped "#" outside a character class and
	//     the next newline character, inclusive, are also ignored. This can
	//     be changed within a pattern by a "(?x)" option setting.
	GRegexExtendedValue RegexCompileFlags = 8
	// The pattern is forced to be "anchored", that is,
	//     it is constrained to match only at the first matching point in the
	//     string that is being searched. This effect can also be achieved by
	//     appropriate constructs in the pattern itself such as the "^"
	//     metacharacter.
	GRegexAnchoredValue RegexCompileFlags = 16
	// A dollar metacharacter ("$") in the pattern
	//     matches only at the end of the string. Without this option, a
	//     dollar also matches immediately before the final character if
	//     it is a newline (but not before any other newlines). This option
	//     is ignored if %G_REGEX_MULTILINE is set.
	GRegexDollarEndonlyValue RegexCompileFlags = 32
	// Inverts the "greediness" of the quantifiers so that
	//     they are not greedy by default, but become greedy if followed by "?".
	//     It can also be set by a "(?U)" option setting within the pattern.
	GRegexUngreedyValue RegexCompileFlags = 512
	// Usually strings must be valid UTF-8 strings, using this
	//     flag they are considered as a raw sequence of bytes.
	GRegexRawValue RegexCompileFlags = 2048
	// Disables the use of numbered capturing
	//     parentheses in the pattern. Any opening parenthesis that is not
	//     followed by "?" behaves as if it were followed by "?:" but named
	//     parentheses can still be used for capturing (and they acquire numbers
	//     in the usual way).
	GRegexNoAutoCaptureValue RegexCompileFlags = 4096
	// Optimize the regular expression. If the pattern will
	//     be used many times, then it may be worth the effort to optimize it
	//     to improve the speed of matches.
	GRegexOptimizeValue RegexCompileFlags = 8192
	// Limits an unanchored pattern to match before (or at) the
	//     first newline. Since: 2.34
	GRegexFirstlineValue RegexCompileFlags = 262144
	// Names used to identify capturing subpatterns need not
	//     be unique. This can be helpful for certain types of pattern when it
	//     is known that only one instance of the named subpattern can ever be
	//     matched.
	GRegexDupnamesValue RegexCompileFlags = 524288
	// Usually any newline character or character sequence is
	//     recognized. If this option is set, the only recognized newline character
	//     is '\r'.
	GRegexNewlineCrValue RegexCompileFlags = 1048576
	// Usually any newline character or character sequence is
	//     recognized. If this option is set, the only recognized newline character
	//     is '\n'.
	GRegexNewlineLfValue RegexCompileFlags = 2097152
	// Usually any newline character or character sequence is
	//     recognized. If this option is set, the only recognized newline character
	//     sequence is '\r\n'.
	GRegexNewlineCrlfValue RegexCompileFlags = 3145728
	// Usually any newline character or character sequence
	//     is recognized. If this option is set, the only recognized newline character
	//     sequences are '\r', '\n', and '\r\n'. Since: 2.34
	GRegexNewlineAnycrlfValue RegexCompileFlags = 5242880
	// Usually any newline character or character sequence
	//     is recognised. If this option is set, then "\R" only recognizes the newline
	//    characters '\r', '\n' and '\r\n'. Since: 2.34
	GRegexBsrAnycrlfValue RegexCompileFlags = 8388608
	// Changes behaviour so that it is compatible with
	//     JavaScript rather than PCRE. Since: 2.34
	GRegexJavascriptCompatValue RegexCompileFlags = 33554432
)

// Flags specifying match-time options.
type RegexMatchFlags int

const (

	// The pattern is forced to be "anchored", that is,
	//     it is constrained to match only at the first matching point in the
	//     string that is being searched. This effect can also be achieved by
	//     appropriate constructs in the pattern itself such as the "^"
	//     metacharacter.
	GRegexMatchAnchoredValue RegexMatchFlags = 16
	// Specifies that first character of the string is
	//     not the beginning of a line, so the circumflex metacharacter should
	//     not match before it. Setting this without %G_REGEX_MULTILINE (at
	//     compile time) causes circumflex never to match. This option affects
	//     only the behaviour of the circumflex metacharacter, it does not
	//     affect "\A".
	GRegexMatchNotbolValue RegexMatchFlags = 128
	// Specifies that the end of the subject string is
	//     not the end of a line, so the dollar metacharacter should not match
	//     it nor (except in multiline mode) a newline immediately before it.
	//     Setting this without %G_REGEX_MULTILINE (at compile time) causes
	//     dollar never to match. This option affects only the behaviour of
	//     the dollar metacharacter, it does not affect "\Z" or "\z".
	GRegexMatchNoteolValue RegexMatchFlags = 256
	// An empty string is not considered to be a valid
	//     match if this option is set. If there are alternatives in the pattern,
	//     they are tried. If all the alternatives match the empty string, the
	//     entire match fails. For example, if the pattern "a?b?" is applied to
	//     a string not beginning with "a" or "b", it matches the empty string
	//     at the start of the string. With this flag set, this match is not
	//     valid, so GRegex searches further into the string for occurrences
	//     of "a" or "b".
	GRegexMatchNotemptyValue RegexMatchFlags = 1024
	// Turns on the partial matching feature, for more
	//     documentation on partial matching see g_match_info_is_partial_match().
	GRegexMatchPartialValue RegexMatchFlags = 32768
	// Overrides the newline definition set when
	//     creating a new #GRegex, setting the '\r' character as line terminator.
	GRegexMatchNewlineCrValue RegexMatchFlags = 1048576
	// Overrides the newline definition set when
	//     creating a new #GRegex, setting the '\n' character as line terminator.
	GRegexMatchNewlineLfValue RegexMatchFlags = 2097152
	// Overrides the newline definition set when
	//     creating a new #GRegex, setting the '\r\n' characters sequence as line terminator.
	GRegexMatchNewlineCrlfValue RegexMatchFlags = 3145728
	// Overrides the newline definition set when
	//     creating a new #GRegex, any Unicode newline sequence
	//     is recognised as a newline. These are '\r', '\n' and '\rn', and the
	//     single characters U+000B LINE TABULATION, U+000C FORM FEED (FF),
	//     U+0085 NEXT LINE (NEL), U+2028 LINE SEPARATOR and
	//     U+2029 PARAGRAPH SEPARATOR.
	GRegexMatchNewlineAnyValue RegexMatchFlags = 4194304
	// Overrides the newline definition set when
	//     creating a new #GRegex; any '\r', '\n', or '\r\n' character sequence
	//     is recognized as a newline. Since: 2.34
	GRegexMatchNewlineAnycrlfValue RegexMatchFlags = 5242880
	// Overrides the newline definition for "\R" set when
	//     creating a new #GRegex; only '\r', '\n', or '\r\n' character sequences
	//     are recognized as a newline by "\R". Since: 2.34
	GRegexMatchBsrAnycrlfValue RegexMatchFlags = 8388608
	// Overrides the newline definition for "\R" set when
	//     creating a new #GRegex; any Unicode newline character or character sequence
	//     are recognized as a newline by "\R". These are '\r', '\n' and '\rn', and the
	//     single characters U+000B LINE TABULATION, U+000C FORM FEED (FF),
	//     U+0085 NEXT LINE (NEL), U+2028 LINE SEPARATOR and
	//     U+2029 PARAGRAPH SEPARATOR. Since: 2.34
	GRegexMatchBsrAnyValue RegexMatchFlags = 16777216
	// An alias for %G_REGEX_MATCH_PARTIAL. Since: 2.34
	GRegexMatchPartialSoftValue RegexMatchFlags = 32768
	// Turns on the partial matching feature. In contrast to
	//     to %G_REGEX_MATCH_PARTIAL_SOFT, this stops matching as soon as a partial match
	//     is found, without continuing to search for a possible complete match. See
	//     g_match_info_is_partial_match() for more information. Since: 2.34
	GRegexMatchPartialHardValue RegexMatchFlags = 134217728
	// Like %G_REGEX_MATCH_NOTEMPTY, but only applied to
	//     the start of the matched string. For anchored
	//     patterns this can only happen for pattern containing "\K". Since: 2.34
	GRegexMatchNotemptyAtstartValue RegexMatchFlags = 268435456
)

// Error codes returned by regular expressions functions.
type RegexError int

const (

	// Compilation of the regular expression failed.
	GRegexErrorCompileValue RegexError = 0
	// Optimization of the regular expression failed.
	GRegexErrorOptimizeValue RegexError = 1
	// Replacement failed due to an ill-formed replacement
	//     string.
	GRegexErrorReplaceValue RegexError = 2
	// The match process failed.
	GRegexErrorMatchValue RegexError = 3
	// Internal error of the regular expression engine.
	//     Since 2.16
	GRegexErrorInternalValue RegexError = 4
	// "\\" at end of pattern. Since 2.16
	GRegexErrorStrayBackslashValue RegexError = 101
	// "\\c" at end of pattern. Since 2.16
	GRegexErrorMissingControlCharValue RegexError = 102
	// Unrecognized character follows "\\".
	//     Since 2.16
	GRegexErrorUnrecognizedEscapeValue RegexError = 103
	// Numbers out of order in "{}"
	//     quantifier. Since 2.16
	GRegexErrorQuantifiersOutOfOrderValue RegexError = 104
	// Number too big in "{}" quantifier.
	//     Since 2.16
	GRegexErrorQuantifierTooBigValue RegexError = 105
	// Missing terminating "]" for
	//     character class. Since 2.16
	GRegexErrorUnterminatedCharacterClassValue RegexError = 106
	// Invalid escape sequence
	//     in character class. Since 2.16
	GRegexErrorInvalidEscapeInCharacterClassValue RegexError = 107
	// Range out of order in character class.
	//     Since 2.16
	GRegexErrorRangeOutOfOrderValue RegexError = 108
	// Nothing to repeat. Since 2.16
	GRegexErrorNothingToRepeatValue RegexError = 109
	// Unrecognized character after "(?",
	//     "(?&lt;" or "(?P". Since 2.16
	GRegexErrorUnrecognizedCharacterValue RegexError = 112
	// POSIX named classes are
	//     supported only within a class. Since 2.16
	GRegexErrorPosixNamedClassOutsideClassValue RegexError = 113
	// Missing terminating ")" or ")"
	//     without opening "(". Since 2.16
	GRegexErrorUnmatchedParenthesisValue RegexError = 114
	// Reference to non-existent
	//     subpattern. Since 2.16
	GRegexErrorInexistentSubpatternReferenceValue RegexError = 115
	// Missing terminating ")" after comment.
	//     Since 2.16
	GRegexErrorUnterminatedCommentValue RegexError = 118
	// Regular expression too large.
	//     Since 2.16
	GRegexErrorExpressionTooLargeValue RegexError = 120
	// Failed to get memory. Since 2.16
	GRegexErrorMemoryErrorValue RegexError = 121
	// Lookbehind assertion is not
	//     fixed length. Since 2.16
	GRegexErrorVariableLengthLookbehindValue RegexError = 125
	// Malformed number or name after "(?(".
	//     Since 2.16
	GRegexErrorMalformedConditionValue RegexError = 126
	// Conditional group contains
	//     more than two branches. Since 2.16
	GRegexErrorTooManyConditionalBranchesValue RegexError = 127
	// Assertion expected after "(?(".
	//     Since 2.16
	GRegexErrorAssertionExpectedValue RegexError = 128
	// Unknown POSIX class name.
	//     Since 2.16
	GRegexErrorUnknownPosixClassNameValue RegexError = 130
	// POSIX collating
	//     elements are not supported. Since 2.16
	GRegexErrorPosixCollatingElementsNotSupportedValue RegexError = 131
	// Character value in "\\x{...}" sequence
	//     is too large. Since 2.16
	GRegexErrorHexCodeTooLargeValue RegexError = 134
	// Invalid condition "(?(0)". Since 2.16
	GRegexErrorInvalidConditionValue RegexError = 135
	// \\C not allowed in
	//     lookbehind assertion. Since 2.16
	GRegexErrorSingleByteMatchInLookbehindValue RegexError = 136
	// Recursive call could loop indefinitely.
	//     Since 2.16
	GRegexErrorInfiniteLoopValue RegexError = 140
	// Missing terminator
	//     in subpattern name. Since 2.16
	GRegexErrorMissingSubpatternNameTerminatorValue RegexError = 142
	// Two named subpatterns have
	//     the same name. Since 2.16
	GRegexErrorDuplicateSubpatternNameValue RegexError = 143
	// Malformed "\\P" or "\\p" sequence.
	//     Since 2.16
	GRegexErrorMalformedPropertyValue RegexError = 146
	// Unknown property name after "\\P" or
	//     "\\p". Since 2.16
	GRegexErrorUnknownPropertyValue RegexError = 147
	// Subpattern name is too long
	//     (maximum 32 characters). Since 2.16
	GRegexErrorSubpatternNameTooLongValue RegexError = 148
	// Too many named subpatterns (maximum
	//     10,000). Since 2.16
	GRegexErrorTooManySubpatternsValue RegexError = 149
	// Octal value is greater than "\\377".
	//     Since 2.16
	GRegexErrorInvalidOctalValueValue RegexError = 151
	// "DEFINE" group contains more
	//     than one branch. Since 2.16
	GRegexErrorTooManyBranchesInDefineValue RegexError = 154
	// Repeating a "DEFINE" group is not allowed.
	//     This error is never raised. Since: 2.16 Deprecated: 2.34
	GRegexErrorDefineRepetionValue RegexError = 155
	// Inconsistent newline options.
	//     Since 2.16
	GRegexErrorInconsistentNewlineOptionsValue RegexError = 156
	// "\\g" is not followed by a braced,
	//      angle-bracketed, or quoted name or number, or by a plain number. Since: 2.16
	GRegexErrorMissingBackReferenceValue RegexError = 157
	// relative reference must not be zero. Since: 2.34
	GRegexErrorInvalidRelativeReferenceValue RegexError = 158
	// the backtracing
	//     control verb used does not allow an argument. Since: 2.34
	GRegexErrorBacktrackingControlVerbArgumentForbiddenValue RegexError = 159
	// unknown backtracing
	//     control verb. Since: 2.34
	GRegexErrorUnknownBacktrackingControlVerbValue RegexError = 160
	// number is too big in escape sequence. Since: 2.34
	GRegexErrorNumberTooBigValue RegexError = 161
	// Missing subpattern name. Since: 2.34
	GRegexErrorMissingSubpatternNameValue RegexError = 162
	// Missing digit. Since 2.34
	GRegexErrorMissingDigitValue RegexError = 163
	// In JavaScript compatibility mode,
	//     "[" is an invalid data character. Since: 2.34
	GRegexErrorInvalidDataCharacterValue RegexError = 164
	// different names for subpatterns of the
	//     same number are not allowed. Since: 2.34
	GRegexErrorExtraSubpatternNameValue RegexError = 165
	// the backtracing control
	//     verb requires an argument. Since: 2.34
	GRegexErrorBacktrackingControlVerbArgumentRequiredValue RegexError = 166
	// "\\c" must be followed by an ASCII
	//     character. Since: 2.34
	GRegexErrorInvalidControlCharValue RegexError = 168
	// "\\k" is not followed by a braced, angle-bracketed, or
	//     quoted name. Since: 2.34
	GRegexErrorMissingNameValue RegexError = 169
	// "\\N" is not supported in a class. Since: 2.34
	GRegexErrorNotSupportedInClassValue RegexError = 171
	// too many forward references. Since: 2.34
	GRegexErrorTooManyForwardReferencesValue RegexError = 172
	// the name is too long in "(*MARK)", "(*PRUNE)",
	//     "(*SKIP)", or "(*THEN)". Since: 2.34
	GRegexErrorNameTooLongValue RegexError = 175
	// the character value in the \\u sequence is
	//     too large. Since: 2.34
	GRegexErrorCharacterValueTooLargeValue RegexError = 176
)

var xRegexCheckReplacement func(string, bool, **Error) bool

// Checks whether @replacement is a valid replacement string
// (see g_regex_replace()), i.e. that all escape sequences in
// it are valid.
//
// If @has_references is not %NULL then @replacement is checked
// for pattern references. For instance, replacement text 'foo\n'
// does not contain references and may be evaluated without information
// about actual match, but '\0\1' (whole match followed by first
// subpattern) requires valid #GMatchInfo object.
func RegexCheckReplacement(ReplacementVar string, HasReferencesVar bool) (bool, error) {
	var cerr *Error

	cret := xRegexCheckReplacement(ReplacementVar, HasReferencesVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xRegexEscapeNul func(string, int) string

// Escapes the nul characters in @string to "\x00".  It can be used
// to compile a regex with embedded nul characters.
//
// For completeness, @length can be -1 for a nul-terminated string.
// In this case the output string will be of course equal to @string.
func RegexEscapeNul(StringVar string, LengthVar int) string {

	cret := xRegexEscapeNul(StringVar, LengthVar)
	return cret
}

var xRegexEscapeString func(uintptr, int) string

// Escapes the special characters used for regular expressions
// in @string, for instance "a.b*c" becomes "a\.b\*c". This
// function is useful to dynamically generate regular expressions.
//
// @string can contain nul characters that are replaced with "\0",
// in this case remember to specify the correct length of @string
// in @length.
func RegexEscapeString(StringVar uintptr, LengthVar int) string {

	cret := xRegexEscapeString(StringVar, LengthVar)
	return cret
}

var xRegexMatchSimple func(string, string, RegexCompileFlags, RegexMatchFlags) bool

// Scans for a match in @string for @pattern.
//
// This function is equivalent to g_regex_match() but it does not
// require to compile the pattern with g_regex_new(), avoiding some
// lines of code when you need just to do a match without extracting
// substrings, capture counts, and so on.
//
// If this function is to be called on the same @pattern more than
// once, it's more efficient to compile the pattern once with
// g_regex_new() and then use g_regex_match().
func RegexMatchSimple(PatternVar string, StringVar string, CompileOptionsVar RegexCompileFlags, MatchOptionsVar RegexMatchFlags) bool {

	cret := xRegexMatchSimple(PatternVar, StringVar, CompileOptionsVar, MatchOptionsVar)
	return cret
}

var xRegexSplitSimple func(string, string, RegexCompileFlags, RegexMatchFlags) uintptr

// Breaks the string on the pattern, and returns an array of
// the tokens. If the pattern contains capturing parentheses,
// then the text for each of the substrings will also be returned.
// If the pattern does not match anywhere in the string, then the
// whole string is returned as the first token.
//
// This function is equivalent to g_regex_split() but it does
// not require to compile the pattern with g_regex_new(), avoiding
// some lines of code when you need just to do a split without
// extracting substrings, capture counts, and so on.
//
// If this function is to be called on the same @pattern more than
// once, it's more efficient to compile the pattern once with
// g_regex_new() and then use g_regex_split().
//
// As a special case, the result of splitting the empty string ""
// is an empty vector, not a vector containing a single string.
// The reason for this special case is that being able to represent
// an empty vector is typically more useful than consistent handling
// of empty elements. If you do need to represent empty elements,
// you'll need to check for the empty string before calling this
// function.
//
// A pattern that can match empty strings splits @string into
// separate characters wherever it matches the empty string between
// characters. For example splitting "ab c" using as a separator
// "\s*", you will get "a", "b" and "c".
func RegexSplitSimple(PatternVar string, StringVar string, CompileOptionsVar RegexCompileFlags, MatchOptionsVar RegexMatchFlags) uintptr {

	cret := xRegexSplitSimple(PatternVar, StringVar, CompileOptionsVar, MatchOptionsVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xRegexCheckReplacement, lib, "g_regex_check_replacement")
	core.PuregoSafeRegister(&xRegexEscapeNul, lib, "g_regex_escape_nul")
	core.PuregoSafeRegister(&xRegexEscapeString, lib, "g_regex_escape_string")
	core.PuregoSafeRegister(&xRegexMatchSimple, lib, "g_regex_match_simple")
	core.PuregoSafeRegister(&xRegexSplitSimple, lib, "g_regex_split_simple")

}
