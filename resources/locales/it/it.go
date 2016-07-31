package it

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type it struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'it' locale
func New() locales.Translator {
	return &it{
		locale:  "it",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *it) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'it'
func (t *it) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'it'
func (t *it) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
