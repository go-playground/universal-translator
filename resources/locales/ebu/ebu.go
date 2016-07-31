package ebu

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ebu struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ebu' locale
func New() locales.Translator {
	return &ebu{
		locale:  "ebu",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *ebu) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ebu'
func (t *ebu) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ebu'
func (t *ebu) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
