package sw

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sw struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sw' locale
func New() locales.Translator {
	return &sw{
		locale:  "sw",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sw) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sw'
func (t *sw) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sw'
func (t *sw) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
