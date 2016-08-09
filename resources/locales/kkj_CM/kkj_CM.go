package kkj_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type kkj_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kkj_CM' locale
func New() locales.Translator {
	return &kkj_CM{
		locale:  "kkj_CM",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *kkj_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kkj_CM'
func (t *kkj_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kkj_CM'
func (t *kkj_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
