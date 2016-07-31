package kam

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kam struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kam' locale
func New() locales.Translator {
	return &kam{
		locale:  "kam",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *kam) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kam'
func (t *kam) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'kam'
func (t *kam) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
