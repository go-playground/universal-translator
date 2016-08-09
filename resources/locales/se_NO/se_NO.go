package se_NO

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type se_NO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'se_NO' locale
func New() locales.Translator {
	return &se_NO{
		locale:  "se_NO",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *se_NO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'se_NO'
func (t *se_NO) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'se_NO'
func (t *se_NO) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}
