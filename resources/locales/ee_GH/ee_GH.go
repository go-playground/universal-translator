package ee_GH

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ee_GH struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ee_GH' locale
func New() locales.Translator {
	return &ee_GH{
		locale:   "ee_GH",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *ee_GH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ee_GH'
func (t *ee_GH) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ee_GH'
func (t *ee_GH) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
