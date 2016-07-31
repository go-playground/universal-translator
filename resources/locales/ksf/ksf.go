package ksf

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ksf struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ksf' locale
func New() locales.Translator {
	return &ksf{
		locale:  "ksf",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *ksf) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ksf'
func (t *ksf) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ksf'
func (t *ksf) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
