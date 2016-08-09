package lkt

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lkt struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lkt' locale
func New() locales.Translator {
	return &lkt{
		locale:  "lkt",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *lkt) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lkt'
func (t *lkt) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lkt'
func (t *lkt) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
