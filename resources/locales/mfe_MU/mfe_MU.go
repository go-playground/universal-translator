package mfe_MU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mfe_MU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mfe_MU' locale
func New() locales.Translator {
	return &mfe_MU{
		locale:  "mfe_MU",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mfe_MU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mfe_MU'
func (t *mfe_MU) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mfe_MU'
func (t *mfe_MU) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
