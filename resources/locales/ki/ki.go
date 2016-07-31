package ki

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ki struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ki' locale
func New() locales.Translator {
	return &ki{
		locale:  "ki",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *ki) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ki'
func (t *ki) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ki'
func (t *ki) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
