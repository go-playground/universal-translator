package en_KI

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type en_KI struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'en_KI' locale
func New() locales.Translator {
	return &en_KI{
		locale:  "en_KI",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *en_KI) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_KI'
func (t *en_KI) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'en_KI'
func (t *en_KI) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
