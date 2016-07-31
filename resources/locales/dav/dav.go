package dav

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dav struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dav' locale
func New() locales.Translator {
	return &dav{
		locale:  "dav",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *dav) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dav'
func (t *dav) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'dav'
func (t *dav) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
