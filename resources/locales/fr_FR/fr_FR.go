package fr_FR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_FR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_FR' locale
func New() locales.Translator {
	return &fr_FR{
		locale:  "fr_FR",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_FR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_FR'
func (t *fr_FR) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'fr_FR'
func (t *fr_FR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
