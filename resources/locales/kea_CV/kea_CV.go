package kea_CV

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type kea_CV struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kea_CV' locale
func New() locales.Translator {
	return &kea_CV{
		locale:  "kea_CV",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *kea_CV) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kea_CV'
func (t *kea_CV) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kea_CV'
func (t *kea_CV) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
