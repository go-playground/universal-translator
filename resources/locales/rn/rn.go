package rn

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type rn struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'rn' locale
func New() locales.Translator {
	return &rn{
		locale:  "rn",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *rn) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'rn'
func (t *rn) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'rn'
func (t *rn) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
