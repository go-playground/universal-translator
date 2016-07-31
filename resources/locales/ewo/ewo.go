package ewo

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ewo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ewo' locale
func New() locales.Translator {
	return &ewo{
		locale:  "ewo",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *ewo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ewo'
func (t *ewo) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ewo'
func (t *ewo) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
