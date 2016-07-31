package si_LK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type si_LK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'si_LK' locale
func New() locales.Translator {
	return &si_LK{
		locale:  "si_LK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *si_LK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'si_LK'
func (t *si_LK) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'si_LK'
func (t *si_LK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (n == 0 || n == 1) || (i == 0 && f == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
