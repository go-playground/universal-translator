package kok

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kok struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kok' locale
func New() locales.Translator {
	return &kok{
		locale:  "kok",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *kok) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kok'
func (t *kok) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'kok'
func (t *kok) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
