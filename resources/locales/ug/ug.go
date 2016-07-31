package ug

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ug struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ug' locale
func New() locales.Translator {
	return &ug{
		locale:  "ug",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ug) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ug'
func (t *ug) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ug'
func (t *ug) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
