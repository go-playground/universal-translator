package el_GR

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type el_GR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'el_GR' locale
func New() locales.Translator {
	return &el_GR{
		locale:  "el_GR",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *el_GR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'el_GR'
func (t *el_GR) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'el_GR'
func (t *el_GR) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
