package so_KE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type so_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'so_KE' locale
func New() locales.Translator {
	return &so_KE{
		locale:  "so_KE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *so_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'so_KE'
func (t *so_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'so_KE'
func (t *so_KE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
