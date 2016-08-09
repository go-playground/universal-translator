package pt_PT

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type pt_PT struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'pt_PT' locale
func New() locales.Translator {
	return &pt_PT{
		locale:  "pt_PT",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *pt_PT) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'pt_PT'
func (t *pt_PT) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'pt_PT'
func (t *pt_PT) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 2 && n != 2 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
