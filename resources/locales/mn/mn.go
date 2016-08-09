package mn

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mn struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mn' locale
func New() locales.Translator {
	return &mn{
		locale:  "mn",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *mn) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mn'
func (t *mn) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mn'
func (t *mn) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
