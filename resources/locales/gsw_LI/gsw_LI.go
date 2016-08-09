package gsw_LI

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type gsw_LI struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gsw_LI' locale
func New() locales.Translator {
	return &gsw_LI{
		locale:  "gsw_LI",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *gsw_LI) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gsw_LI'
func (t *gsw_LI) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'gsw_LI'
func (t *gsw_LI) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
