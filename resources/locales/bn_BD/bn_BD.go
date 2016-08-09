package bn_BD

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type bn_BD struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bn_BD' locale
func New() locales.Translator {
	return &bn_BD{
		locale:  "bn_BD",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *bn_BD) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bn_BD'
func (t *bn_BD) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bn_BD'
func (t *bn_BD) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
