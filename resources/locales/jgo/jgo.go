package jgo

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type jgo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'jgo' locale
func New() locales.Translator {
	return &jgo{
		locale:  "jgo",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *jgo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'jgo'
func (t *jgo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'jgo'
func (t *jgo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
