package kok_IN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kok_IN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kok_IN' locale
func New() locales.Translator {
	return &kok_IN{
		locale:  "kok_IN",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *kok_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kok_IN'
func (t *kok_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'kok_IN'
func (t *kok_IN) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
