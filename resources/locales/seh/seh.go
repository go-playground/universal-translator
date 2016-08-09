package seh

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type seh struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'seh' locale
func New() locales.Translator {
	return &seh{
		locale:  "seh",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *seh) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'seh'
func (t *seh) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'seh'
func (t *seh) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
