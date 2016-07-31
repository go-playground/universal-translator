package nmg

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nmg struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nmg' locale
func New() locales.Translator {
	return &nmg{
		locale:  "nmg",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *nmg) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nmg'
func (t *nmg) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'nmg'
func (t *nmg) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
