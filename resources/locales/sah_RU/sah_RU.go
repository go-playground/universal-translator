package sah_RU

import (
	"math"

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

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sah_RU'
func (t *sah_RU) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
