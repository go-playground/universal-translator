package mer_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mer_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mer_KE' locale
func New() locales.Translator {
	return &mer_KE{
		locale:  "mer_KE",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mer_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mer_KE'
func (t *mer_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mer_KE'
func (t *mer_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
