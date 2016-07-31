package lrc_IQ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lrc_IQ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lrc_IQ' locale
func New() locales.Translator {
	return &lrc_IQ{
		locale:  "lrc_IQ",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *lrc_IQ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lrc_IQ'
func (t *lrc_IQ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'lrc_IQ'
func (t *lrc_IQ) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
