package de_LU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type de_LU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'de_LU' locale
func New() locales.Translator {
	return &de_LU{
		locale:  "de_LU",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *de_LU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'de_LU'
func (t *de_LU) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'de_LU'
func (t *de_LU) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
