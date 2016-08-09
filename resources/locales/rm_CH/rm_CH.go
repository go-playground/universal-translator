package rm_CH

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type rm_CH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'rm_CH' locale
func New() locales.Translator {
	return &rm_CH{
		locale:  "rm_CH",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *rm_CH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'rm_CH'
func (t *rm_CH) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'rm_CH'
func (t *rm_CH) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
