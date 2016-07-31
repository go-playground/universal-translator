package lu

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lu struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lu' locale
func New() locales.Translator {
	return &lu{
		locale:  "lu",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *lu) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lu'
func (t *lu) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'lu'
func (t *lu) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
