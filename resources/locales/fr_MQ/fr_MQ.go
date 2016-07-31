package fr_MQ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_MQ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_MQ' locale
func New() locales.Translator {
	return &fr_MQ{
		locale:  "fr_MQ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_MQ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_MQ'
func (t *fr_MQ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'fr_MQ'
func (t *fr_MQ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
