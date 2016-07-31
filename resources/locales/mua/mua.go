package mua

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mua struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mua' locale
func New() locales.Translator {
	return &mua{
		locale:  "mua",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mua) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mua'
func (t *mua) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mua'
func (t *mua) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
