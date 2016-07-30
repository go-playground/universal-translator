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

// Translator encapsulates an instance of
type Translator interface {
	Locale() string
}

// PluralStringToInt returns the enum value of 'plural' provided
func PluralStringToInt(plural string) PluralRule {

	switch plural {
	case "zero":
		return PluralRuleZero
	case "one":
		return PluralRuleOne
	case "two":
		return PluralRuleTwo
	case "few":
		return PluralRuleFew
	case "many":
		return PluralRuleMany
	case "other":
		return PluralRuleOther
	default:
		return PluralRuleUnknown
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
