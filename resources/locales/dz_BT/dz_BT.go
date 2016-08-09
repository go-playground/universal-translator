package dz_BT

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type dz_BT struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dz_BT' locale
func New() locales.Translator {
	return &dz_BT{
		locale:  "dz_BT",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *dz_BT) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dz_BT'
func (t *dz_BT) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'dz_BT'
func (t *dz_BT) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
