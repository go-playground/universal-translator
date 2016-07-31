package luo

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type luo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'luo' locale
func New() locales.Translator {
	return &luo{
		locale:  "luo",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *luo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'luo'
func (t *luo) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'luo'
func (t *luo) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
