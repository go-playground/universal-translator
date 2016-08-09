package ug_CN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ug_CN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ug_CN' locale
func New() locales.Translator {
	return &ug_CN{
		locale:  "ug_CN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ug_CN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ug_CN'
func (t *ug_CN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ug_CN'
func (t *ug_CN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
