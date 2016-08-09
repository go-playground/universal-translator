package yo

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type yo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'yo' locale
func New() locales.Translator {
	return &yo{
		locale:  "yo",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *yo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'yo'
func (t *yo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'yo'
func (t *yo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
