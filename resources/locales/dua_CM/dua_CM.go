package dua_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dua_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dua_CM' locale
func New() locales.Translator {
	return &dua_CM{
		locale:  "dua_CM",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *dua_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dua_CM'
func (t *dua_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'dua_CM'
func (t *dua_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
