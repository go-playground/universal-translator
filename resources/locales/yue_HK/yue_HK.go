package yue_HK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type yue_HK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'yue_HK' locale
func New() locales.Translator {
	return &yue_HK{
		locale:  "yue_HK",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *yue_HK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'yue_HK'
func (t *yue_HK) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'yue_HK'
func (t *yue_HK) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
