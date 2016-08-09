package en_SS

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type en_SS struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'en_SS' locale
func New() locales.Translator {
	return &en_SS{
		locale:  "en_SS",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *en_SS) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_SS'
func (t *en_SS) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'en_SS'
func (t *en_SS) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
