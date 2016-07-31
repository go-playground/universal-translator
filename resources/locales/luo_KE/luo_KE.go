package luo_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type luo_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'luo_KE' locale
func New() locales.Translator {
	return &luo_KE{
		locale:  "luo_KE",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *luo_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'luo_KE'
func (t *luo_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'luo_KE'
func (t *luo_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
