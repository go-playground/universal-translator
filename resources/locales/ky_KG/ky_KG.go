package ky_KG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ky_KG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ky_KG' locale
func New() locales.Translator {
	return &ky_KG{
		locale:  "ky_KG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ky_KG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ky_KG'
func (t *ky_KG) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ky_KG'
func (t *ky_KG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
