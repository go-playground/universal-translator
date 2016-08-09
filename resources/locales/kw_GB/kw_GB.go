package kw_GB

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type kw_GB struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kw_GB' locale
func New() locales.Translator {
	return &kw_GB{
		locale:  "kw_GB",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *kw_GB) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kw_GB'
func (t *kw_GB) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kw_GB'
func (t *kw_GB) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}
