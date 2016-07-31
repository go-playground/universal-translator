package mfe

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mfe struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mfe' locale
func New() locales.Translator {
	return &mfe{
		locale:  "mfe",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mfe) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mfe'
func (t *mfe) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mfe'
func (t *mfe) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
