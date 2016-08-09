package brx

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type brx struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'brx' locale
func New() locales.Translator {
	return &brx{
		locale:  "brx",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *brx) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'brx'
func (t *brx) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'brx'
func (t *brx) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
