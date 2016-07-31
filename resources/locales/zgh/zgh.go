package zgh

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zgh struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zgh' locale
func New() locales.Translator {
	return &zgh{
		locale:  "zgh",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *zgh) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zgh'
func (t *zgh) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'zgh'
func (t *zgh) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
