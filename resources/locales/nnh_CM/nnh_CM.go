package nnh_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nnh_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nnh_CM' locale
func New() locales.Translator {
	return &nnh_CM{
		locale:  "nnh_CM",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nnh_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nnh_CM'
func (t *nnh_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'nnh_CM'
func (t *nnh_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
