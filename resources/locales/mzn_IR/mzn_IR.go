package mzn_IR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mzn_IR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mzn_IR' locale
func New() locales.Translator {
	return &mzn_IR{
		locale:  "mzn_IR",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mzn_IR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mzn_IR'
func (t *mzn_IR) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mzn_IR'
func (t *mzn_IR) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
