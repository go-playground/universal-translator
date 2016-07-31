package de_LI

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type de_LI struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'de_LI' locale
func New() locales.Translator {
	return &de_LI{
		locale:  "de_LI",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *de_LI) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'de_LI'
func (t *de_LI) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'de_LI'
func (t *de_LI) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
