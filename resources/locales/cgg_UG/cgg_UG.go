package cgg_UG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type cgg_UG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'cgg_UG' locale
func New() locales.Translator {
	return &cgg_UG{
		locale:  "cgg_UG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *cgg_UG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'cgg_UG'
func (t *cgg_UG) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'cgg_UG'
func (t *cgg_UG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
