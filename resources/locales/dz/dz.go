package dz

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dz struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dz' locale
func New() locales.Translator {
	return &dz{
		locale:  "dz",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *dz) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dz'
func (t *dz) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'dz'
func (t *dz) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
