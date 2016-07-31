package kam_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kam_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kam_KE' locale
func New() locales.Translator {
	return &kam_KE{
		locale:  "kam_KE",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *kam_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kam_KE'
func (t *kam_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'kam_KE'
func (t *kam_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
