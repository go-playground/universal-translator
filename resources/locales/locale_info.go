package locales

import (
	"fmt"
	"strconv"
	"strings"
)

// ErrBadNumberValue is returned when the number passed for
// plural rule determination cannot be parsed
type ErrBadNumberValue struct {
	NumberValue string
	InnerError  error
}

// Error returns ErrBadNumberValue error string
func (e *ErrBadNumberValue) Error() string {
	return fmt.Sprintf("Invalid Number Value '%s' %s", e.NumberValue, e.InnerError)
}

var _ error = new(ErrBadNumberValue)

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

	// CardinalPluralRule returns the plural rule needed
	// determined by the provided 'num' variable
	CardinalPluralRule(num string) (PluralRule, error)
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

// N returns the absolute value of the source number (integer and decimals).
func N(num string) (float64, error) {

	if num[0] == '-' {
		num = num[1:]
	}

	f, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}

// I returns the integer digits of N.
func I(num string) (int64, error) {

	i, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// V returns the number of visible fraction digits in N, with trailing zeros.
func V(num string) (v int64) {

	idx := strings.Index(num, ".")

	if idx != -1 {
		v = int64(len(num[idx+1:]))
	}

	return
}

// W returns the number of visible fraction digits in N, without trailing zeros.
func W(num string) (w int64, err error) {

	idx := strings.Index(num, ".")

	if idx != -1 {

		idx++
		end := len(num[idx:]) + 1

		for i := end; i >= 0; i-- {
			if num[i] != '0' {
				end = i + 1
				break
			}
		}

		w = int64(len(num[idx:end]))
	}

	return w, nil
}

// F returns the visible fractional digits in N, with trailing zeros.
func F(num string) (w int64, err error) {

	idx := strings.Index(num, ".")

	if idx != -1 {

		w, err = strconv.ParseInt(num[idx+1:], 10, 64)
		if err != nil {
			return
		}
	}

	return
}

// T returns the visible fractional digits in N, without trailing zeros.
func T(num string) (t int64, err error) {

	idx := strings.Index(num, ".")

	if idx != -1 {

		idx++
		end := len(num[idx:]) + 1

		for i := end; i >= 0; i-- {
			if num[i] != '0' {
				end = i + 1
				break
			}
		}

		t, err = strconv.ParseInt(num[idx:end], 10, 64)
		if err != nil {
			return
		}

		t = int64(t)
	}

	return t, nil
}
