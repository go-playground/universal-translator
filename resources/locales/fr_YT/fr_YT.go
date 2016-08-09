package fr_YT

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_YT struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fr_YT' locale
func New() locales.Translator {
	return &fr_YT{
		locale:  "fr_YT",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fr_YT) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fr_YT'
func (t *fr_YT) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fr_YT'
func (t *fr_YT) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
