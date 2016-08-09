package gsw_CH

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type gsw_CH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gsw_CH' locale
func New() locales.Translator {
	return &gsw_CH{
		locale:  "gsw_CH",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *gsw_CH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gsw_CH'
func (t *gsw_CH) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'gsw_CH'
func (t *gsw_CH) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
