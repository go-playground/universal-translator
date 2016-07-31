package fr_TD

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_TD struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_TD' locale
func New() locales.Translator {
	return &fr_TD{
		locale:  "fr_TD",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_TD) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_TD'
func (t *fr_TD) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'fr_TD'
func (t *fr_TD) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
