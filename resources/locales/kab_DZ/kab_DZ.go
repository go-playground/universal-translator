package kab_DZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type kab_DZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kab_DZ' locale
func New() locales.Translator {
	return &kab_DZ{
		locale:  "kab_DZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *kab_DZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kab_DZ'
func (t *kab_DZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kab_DZ'
func (t *kab_DZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
