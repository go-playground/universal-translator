package lrc_IR

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lrc_IR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lrc_IR' locale
func New() locales.Translator {
	return &lrc_IR{
		locale:  "lrc_IR",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *lrc_IR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lrc_IR'
func (t *lrc_IR) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lrc_IR'
func (t *lrc_IR) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
