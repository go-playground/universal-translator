package dyo

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dyo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dyo' locale
func New() locales.Translator {
	return &dyo{
		locale:  "dyo",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *dyo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dyo'
func (t *dyo) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'dyo'
func (t *dyo) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
