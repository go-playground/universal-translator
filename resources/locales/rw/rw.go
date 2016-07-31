package rw

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type rw struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'rw' locale
func New() locales.Translator {
	return &rw{
		locale:  "rw",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *rw) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'rw'
func (t *rw) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'rw'
func (t *rw) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
