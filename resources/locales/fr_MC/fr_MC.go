package fr_MC

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_MC struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_MC' locale
func New() locales.Translator {
	return &fr_MC{
		locale:  "fr_MC",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_MC) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_MC'
func (t *fr_MC) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'fr_MC'
func (t *fr_MC) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
