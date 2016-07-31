package twq_NE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type twq_NE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'twq_NE' locale
func New() locales.Translator {
	return &twq_NE{
		locale:  "twq_NE",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *twq_NE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'twq_NE'
func (t *twq_NE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'twq_NE'
func (t *twq_NE) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
