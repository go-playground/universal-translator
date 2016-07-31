package kln_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kln_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kln_KE' locale
func New() locales.Translator {
	return &kln_KE{
		locale:  "kln_KE",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *kln_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kln_KE'
func (t *kln_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'kln_KE'
func (t *kln_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
