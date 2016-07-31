package guz_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type guz_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'guz_KE' locale
func New() locales.Translator {
	return &guz_KE{
		locale:  "guz_KE",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *guz_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'guz_KE'
func (t *guz_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'guz_KE'
func (t *guz_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
