package guz

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type guz struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'guz' locale
func New() locales.Translator {
	return &guz{
		locale:  "guz",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *guz) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'guz'
func (t *guz) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'guz'
func (t *guz) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
