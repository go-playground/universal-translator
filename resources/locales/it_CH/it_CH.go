package it_CH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type it_CH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'it_CH' locale
func New() locales.Translator {
	return &it_CH{
		locale:  "it_CH",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *it_CH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'it_CH'
func (t *it_CH) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'it_CH'
func (t *it_CH) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
