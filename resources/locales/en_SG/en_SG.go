package en_SG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type en_SG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'en_SG' locale
func New() locales.Translator {
	return &en_SG{
		locale:  "en_SG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *en_SG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_SG'
func (t *en_SG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'en_SG'
func (t *en_SG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
