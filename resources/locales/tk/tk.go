package tk

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type tk struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'tk' locale
func New() locales.Translator {
	return &tk{
		locale:  "tk",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *tk) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'tk'
func (t *tk) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'tk'
func (t *tk) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
