package ckb_IR

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ckb_IR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ckb_IR' locale
func New() locales.Translator {
	return &ckb_IR{
		locale:  "ckb_IR",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ckb_IR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ckb_IR'
func (t *ckb_IR) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ckb_IR'
func (t *ckb_IR) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
