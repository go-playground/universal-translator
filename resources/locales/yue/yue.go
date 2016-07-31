package yue

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type yue struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'yue' locale
func New() locales.Translator {
	return &yue{
		locale:  "yue",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *yue) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'yue'
func (t *yue) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'yue'
func (t *yue) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
