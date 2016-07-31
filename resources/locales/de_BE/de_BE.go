package de_BE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type de_BE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'de_BE' locale
func New() locales.Translator {
	return &de_BE{
		locale:  "de_BE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *de_BE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'de_BE'
func (t *de_BE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'de_BE'
func (t *de_BE) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
