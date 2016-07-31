package saq

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type saq struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'saq' locale
func New() locales.Translator {
	return &saq{
		locale:  "saq",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *saq) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'saq'
func (t *saq) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'saq'
func (t *saq) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
