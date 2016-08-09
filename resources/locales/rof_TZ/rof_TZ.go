package rof_TZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type rof_TZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'rof_TZ' locale
func New() locales.Translator {
	return &rof_TZ{
		locale:  "rof_TZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *rof_TZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'rof_TZ'
func (t *rof_TZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'rof_TZ'
func (t *rof_TZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
