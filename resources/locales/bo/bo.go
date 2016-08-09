package bo

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type bo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bo' locale
func New() locales.Translator {
	return &bo{
		locale:  "bo",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *bo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bo'
func (t *bo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bo'
func (t *bo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
