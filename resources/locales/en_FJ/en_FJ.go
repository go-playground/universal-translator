package en_FJ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type en_FJ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'en_FJ' locale
func New() locales.Translator {
	return &en_FJ{
		locale:  "en_FJ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *en_FJ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_FJ'
func (t *en_FJ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'en_FJ'
func (t *en_FJ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
