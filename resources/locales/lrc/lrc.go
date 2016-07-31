package lrc

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lrc struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lrc' locale
func New() locales.Translator {
	return &lrc{
		locale:  "lrc",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *lrc) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lrc'
func (t *lrc) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'lrc'
func (t *lrc) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
