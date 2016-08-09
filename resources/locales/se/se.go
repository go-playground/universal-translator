package se

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type se struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'se' locale
func New() locales.Translator {
	return &se{
		locale:  "se",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *se) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'se'
func (t *se) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'se'
func (t *se) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}
