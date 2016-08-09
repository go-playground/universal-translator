package se_FI

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type se_FI struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'se_FI' locale
func New() locales.Translator {
	return &se_FI{
		locale:  "se_FI",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *se_FI) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'se_FI'
func (t *se_FI) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'se_FI'
func (t *se_FI) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}
