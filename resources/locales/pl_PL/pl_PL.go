package pl_PL

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type pl_PL struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'pl_PL' locale
func New() locales.Translator {
	return &pl_PL{
		locale:  "pl_PL",
		plurals: []locales.PluralRule{2, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *pl_PL) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'pl_PL'
func (t *pl_PL) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'pl_PL'
func (t *pl_PL) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if v == 0 && i%10 >= 2 && i%10 <= 4 && i%100 < 12 && i%100 > 14 {
		return locales.PluralRuleFew
	} else if (v == 0 && i != 1 && i%10 >= 0 && i%10 <= 1) || (v == 0 && i%10 >= 5 && i%10 <= 9) || (v == 0 && i%100 >= 12 && i%100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
