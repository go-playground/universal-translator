package ce_RU

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ce_RU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ce_RU' locale
func New() locales.Translator {
	return &ce_RU{
		locale:  "ce_RU",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ce_RU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ce_RU'
func (t *ce_RU) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ce_RU'
func (t *ce_RU) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
