package rw_RW

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type rw_RW struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'rw_RW' locale
func New() locales.Translator {
	return &rw_RW{
		locale:  "rw_RW",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *rw_RW) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'rw_RW'
func (t *rw_RW) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'rw_RW'
func (t *rw_RW) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
