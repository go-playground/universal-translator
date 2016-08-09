package cs_CZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type cs_CZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'cs_CZ' locale
func New() locales.Translator {
	return &cs_CZ{
		locale:  "cs_CZ",
		plurals: []locales.PluralRule{2, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *cs_CZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'cs_CZ'
func (t *cs_CZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'cs_CZ'
func (t *cs_CZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if i >= 2 && i <= 4 && v == 0 {
		return locales.PluralRuleFew
	} else if v != 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
