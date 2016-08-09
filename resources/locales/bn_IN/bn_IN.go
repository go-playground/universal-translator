package bn_IN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type bn_IN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bn_IN' locale
func New() locales.Translator {
	return &bn_IN{
		locale:  "bn_IN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *bn_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bn_IN'
func (t *bn_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bn_IN'
func (t *bn_IN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
