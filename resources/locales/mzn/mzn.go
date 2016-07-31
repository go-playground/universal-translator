package mzn

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mzn struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mzn' locale
func New() locales.Translator {
	return &mzn{
		locale:  "mzn",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mzn) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mzn'
func (t *mzn) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mzn'
func (t *mzn) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
