package sah

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sah struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sah' locale
func New() locales.Translator {
	return &sah{
		locale:  "sah",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *sah) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sah'
func (t *sah) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sah'
func (t *sah) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
