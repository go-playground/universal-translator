package pt_MO

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type pt_MO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'pt_MO' locale
func New() locales.Translator {
	return &pt_MO{
		locale:  "pt_MO",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *pt_MO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'pt_MO'
func (t *pt_MO) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'pt_MO'
func (t *pt_MO) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 2 && n != 2 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
