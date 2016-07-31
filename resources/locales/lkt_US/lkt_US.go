package lkt_US

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lkt_US struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lkt_US' locale
func New() locales.Translator {
	return &lkt_US{
		locale:  "lkt_US",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *lkt_US) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lkt_US'
func (t *lkt_US) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'lkt_US'
func (t *lkt_US) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
