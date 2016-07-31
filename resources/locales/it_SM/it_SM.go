package it_SM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type it_SM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'it_SM' locale
func New() locales.Translator {
	return &it_SM{
		locale:  "it_SM",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *it_SM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'it_SM'
func (t *it_SM) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'it_SM'
func (t *it_SM) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
