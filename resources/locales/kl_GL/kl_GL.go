package kl_GL

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type kl_GL struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kl_GL' locale
func New() locales.Translator {
	return &kl_GL{
		locale:  "kl_GL",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *kl_GL) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kl_GL'
func (t *kl_GL) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kl_GL'
func (t *kl_GL) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
