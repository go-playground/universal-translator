package vun

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type vun struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vun' locale
func New() locales.Translator {
	return &vun{
		locale:  "vun",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *vun) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vun'
func (t *vun) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'vun'
func (t *vun) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
