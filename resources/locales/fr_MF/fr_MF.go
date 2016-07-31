package fr_MF

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_MF struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_MF' locale
func New() locales.Translator {
	return &fr_MF{
		locale:  "fr_MF",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_MF) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_MF'
func (t *fr_MF) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'fr_MF'
func (t *fr_MF) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
