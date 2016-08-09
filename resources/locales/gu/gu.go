package gu

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type gu struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gu' locale
func New() locales.Translator {
	return &gu{
		locale:  "gu",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *gu) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gu'
func (t *gu) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'gu'
func (t *gu) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
