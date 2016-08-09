package kw

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type kw struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kw' locale
func New() locales.Translator {
	return &kw{
		locale:  "kw",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *kw) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kw'
func (t *kw) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kw'
func (t *kw) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}
