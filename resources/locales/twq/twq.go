package twq

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type twq struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'twq' locale
func New() locales.Translator {
	return &twq{
		locale:  "twq",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *twq) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'twq'
func (t *twq) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'twq'
func (t *twq) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
