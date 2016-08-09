package gv

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type gv struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gv' locale
func New() locales.Translator {
	return &gv{
		locale:  "gv",
		plurals: []locales.PluralRule{2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *gv) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gv'
func (t *gv) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'gv'
func (t *gv) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if v == 0 && i%10 == 1 {
		return locales.PluralRuleOne
	} else if v == 0 && i%10 == 2 {
		return locales.PluralRuleTwo
	} else if v == 0 && (i%100 == 0 || i%100 == 20 || i%100 == 40 || i%100 == 60 || i%100 == 80) {
		return locales.PluralRuleFew
	} else if v != 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
