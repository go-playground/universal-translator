package nus_SS

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nus_SS struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nus_SS' locale
func New() locales.Translator {
	return &nus_SS{
		locale:  "nus_SS",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *nus_SS) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nus_SS'
func (t *nus_SS) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'nus_SS'
func (t *nus_SS) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
