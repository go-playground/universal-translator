package fr_DZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_DZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_DZ' locale
func New() locales.Translator {
	return &fr_DZ{
		locale:  "fr_DZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_DZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_DZ'
func (t *fr_DZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'fr_DZ'
func (t *fr_DZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
