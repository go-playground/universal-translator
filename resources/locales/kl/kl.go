package kl

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type kl struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kl' locale
func New() locales.Translator {
	return &kl{
		locale:  "kl",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *kl) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kl'
func (t *kl) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kl'
func (t *kl) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
