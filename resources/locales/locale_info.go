package locales

import "strconv"

// // ErrBadNumberValue is returned when the number passed for
// // plural rule determination cannot be parsed
// type ErrBadNumberValue struct {
// 	NumberValue string
// 	InnerError  error
// }

// // Error returns ErrBadNumberValue error string
// func (e *ErrBadNumberValue) Error() string {
// 	return fmt.Sprintf("Invalid Number Value '%s' %s", e.NumberValue, e.InnerError)
// }

// var _ error = new(ErrBadNumberValue)

// PluralRule denotes the type of plural rules
type PluralRule int

// PluralRule's
const (
	PluralRuleUnknown PluralRule = iota
	PluralRuleZero               // zero
	PluralRuleOne                // one - singular
	PluralRuleTwo                // two - dual
	PluralRuleFew                // few - paucal
	PluralRuleMany               // many - also used for fractions if they have a separate class
	PluralRuleOther              // other - required—general plural form—also used if the language only has a single form
)

const (
	pluralsString = "UnknownZeroOneTwoFewManyOther"
)

// Translator encapsulates an instance of
type Translator interface {

	// Locale returns the string value of the translator
	Locale() string

	// Plurals returns an array of plural rules associated
	// with this translator
	Plurals() []PluralRule
}

// String returns the string value  of PluralRule
func (p PluralRule) String() string {

	switch p {
	case PluralRuleZero:
		return pluralsString[7:11]
	case PluralRuleOne:
		return pluralsString[11:14]
	case PluralRuleTwo:
		return pluralsString[14:17]
	case PluralRuleFew:
		return pluralsString[17:20]
	case PluralRuleMany:
		return pluralsString[20:24]
	case PluralRuleOther:
		return pluralsString[24:]
	default:
		return pluralsString[:7]
	}
}

//
// Precision Notes:
//
// must specify a precision >= 0, and here is why https://play.golang.org/p/LyL90U0Vyh
//
// 	v := float64(3.141)
// 	i := float64(int64(v))
//
// 	fmt.Println(v - i)
//
// 	or
//
// 	s := strconv.FormatFloat(v-i, 'f', -1, 64)
// 	fmt.Println(s)
//
// these will not print what you'd expect: 0.14100000000000001
// and so this library requires a precision to be specified, or
// inaccurate plural rules could be applied.
//
//
//
// n - absolute value of the source number (integer and decimals).
// i - integer digits of n.
// v - number of visible fraction digits in n, with trailing zeros.
// w - number of visible fraction digits in n, without trailing zeros.
// f - visible fractional digits in n, with trailing zeros.
// t - visible fractional digits in n, without trailing zeros.
//
//
// Func(num float64, v uint64) // v = digits/precision and prevents -1 as a special case as this can lead to very unexpected behaviour, see precision note's above.
//
// n := math.Abs(num)
// i := int64(n)
// v := v
//
//
// w := strconv.FormatFloat(num-float64(i), 'f', int(v), 64)  // then parse backwards on string until no more zero's....
// f := strconv.FormatFloat(n, 'f', int(v), 64) 			  // then turn everything after decimal into an int64
// t := strconv.FormatFloat(n, 'f', int(v), 64) 			  // then parse backwards on string until no more zero's....
//
//
//
// General Inclusion Rules
// - v will always be available inherently
// - all require n
// - w requires i
//

// W returns the number of visible fraction digits in N, without trailing zeros.
func W(n float64, v uint64) (w int64) {

	s := strconv.FormatFloat(n-float64(int64(n)), 'f', int(v), 64)

	// with either be '0' or '0.xxxx', so if 1 then w will be zero
	// otherwise need to parse
	if len(s) != 1 {

		s = s[2:]
		end := len(s) + 1

		for i := end; i >= 0; i-- {
			if s[i] != '0' {
				end = i + 1
				break
			}
		}

		w = int64(len(s[:end]))
	}

	return
}

// F returns the visible fractional digits in N, with trailing zeros.
func F(n float64, v uint64) (f int64) {

	s := strconv.FormatFloat(n-float64(int64(n)), 'f', int(v), 64)

	// with either be '0' or '0.xxxx', so if 1 then f will be zero
	// otherwise need to parse
	if len(s) != 1 {

		// ignoring error, because it can't fail as we generated
		// the string internally from a real number
		f, _ = strconv.ParseInt(s[2:], 10, 64)
	}

	return
}

// T returns the visible fractional digits in N, without trailing zeros.
func T(n float64, v uint64) (t int64) {

	s := strconv.FormatFloat(n-float64(int64(n)), 'f', int(v), 64)

	// with either be '0' or '0.xxxx', so if 1 then t will be zero
	// otherwise need to parse
	if len(s) != 1 {

		s = s[2:]
		end := len(s) + 1

		for i := end; i >= 0; i-- {
			if s[i] != '0' {
				end = i + 1
				break
			}
		}

		// ignoring error, because it can't fail as we generated
		// the string internally from a real number
		t, _ = strconv.ParseInt(s[:end], 10, 64)
	}

	return
}
