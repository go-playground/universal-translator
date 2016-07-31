package mua_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mua_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mua_CM' locale
func New() locales.Translator {
	return &mua_CM{
		locale:  "mua_CM",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mua_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mua_CM'
func (t *mua_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mua_CM'
func (t *mua_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
