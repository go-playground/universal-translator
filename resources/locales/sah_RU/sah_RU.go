package sah_RU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sah_RU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sah_RU' locale
func New() locales.Translator {
	return &sah_RU{
		locale:  "sah_RU",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *sah_RU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sah_RU'
func (t *sah_RU) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sah_RU'
func (t *sah_RU) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
